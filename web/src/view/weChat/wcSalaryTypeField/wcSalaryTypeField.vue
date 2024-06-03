<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">

        <el-form-item label="职工等级类型">
          <el-select v-model="searchInfo.type" clearable placeholder="请选择工资类型">
            <el-option v-for="type in types" :key="type.value" :label="type.label" :value="type.value"></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
<!--        <div class="gva-btn-list">-->
<!--            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>-->
<!--            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>-->
<!--        </div>-->
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="工资类型" prop="typeText" width="240" />
        <el-table-column align="left" label="款项字段" prop="field" width="200" />
        <el-table-column align="left" label="款项定义" prop="name" width="180" />
        <el-table-column align="left" label="是否必选" prop="isRequiredText" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateWcSalaryTypeFieldFunc(scope.row)">变更</el-button>
<!--            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>-->
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #title>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="工资类型:"  prop="type" >
              <el-select v-model="formData.type" placeholder="请选择工资类型" :disabled="type==='update'?'disabled':false">
                <el-option v-for="type in types" :key="type.value" :label="type.label" :value="type.value"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="款项字段:"  prop="field" >
              <el-input v-model="formData.field" :clearable="true"  placeholder="请输入款项字段" :disabled="type==='update'?'disabled':false"/>
            </el-form-item>
            <el-form-item label="款项定义:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入款项定义" />
            </el-form-item>
            <el-form-item label="是否必选:"  prop="isRequired" >
              <el-select v-model="formData.isRequired" placeholder="选择是否必选">
                <el-option v-for="isRequired in isRequiredS" :key="isRequired.value" :label="isRequired.label" :value="isRequired.value"></el-option>
              </el-select>
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer size="800" v-model="detailShow" :before-close="closeDetailShow" title="查看详情" destroy-on-close>
          <template #title>
             <div class="flex justify-between items-center">
               <span class="text-lg">查看详情</span>
             </div>
         </template>
        <el-descriptions :column="1" border>
                <el-descriptions-item label="工资类型">
                        {{ formData.type }}
                </el-descriptions-item>
                <el-descriptions-item label="款项字段">
                        {{ formData.field }}
                </el-descriptions-item>
                <el-descriptions-item label="款项定义">
                        {{ formData.name }}
                </el-descriptions-item>
                <el-descriptions-item label="是否必选">
                        {{ formData.isRequired }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createWcSalaryTypeField,
  deleteWcSalaryTypeField,
  deleteWcSalaryTypeFieldByIds,
  updateWcSalaryTypeField,
  findWcSalaryTypeField,
  getWcSalaryTypeFieldList
} from '@/api/weChat/wcSalaryTypeField'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import {InfoFilled} from "@element-plus/icons-vue";
import SelectRankType from "@/components/selectRankType/index.vue";

defineOptions({
    name: 'WcSalaryTypeField'
})

const types = ref([
  { label: '基本工资', value: 1 },
  { label: '集团经营绩效奖励', value: 2 },
  { label: '节日金', value: 3 },
  { label: '半年奖', value: 4 },
  { label: '年度奖金', value: 5 },
  { label: '总部职能体系月度奖金', value: 6 },
  { label: '总部金纳斯市场体系月度奖金', value: 7 },
  { label: '总部调理中心体系月度奖金', value: 8 },
])

const isRequiredS = ref([
  { label: '否', value: 0 },
  { label: '是', value: 1 },
])

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        type: '',
        field: '',
        name: '',
        isRequired: 1,
        typeText: '',
        isRequiredText: '',
        })


// 验证规则
const rule = reactive({
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               field : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               isRequired : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getWcSalaryTypeFieldList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteWcSalaryTypeFieldFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteWcSalaryTypeFieldByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateWcSalaryTypeFieldFunc = async(row) => {
    const res = await findWcSalaryTypeField({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rewcSalaryTypeField
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWcSalaryTypeFieldFunc = async (row) => {
    const res = await deleteWcSalaryTypeField({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findWcSalaryTypeField({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.rewcSalaryTypeField
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          type: '',
          field: '',
          name: '',
          isRequired: 1,
          typeText: '',
          isRequiredText: '',
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        type: '',
        field: '',
        name: '',
        isRequired: 1,
        typeText: '',
        isRequiredText: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createWcSalaryTypeField(formData.value)
                  break
                case 'update':
                  res = await updateWcSalaryTypeField(formData.value)
                  break
                default:
                  res = await createWcSalaryTypeField(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
