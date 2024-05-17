<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="成员名称" prop="staffName" width="120"/>
        <el-table-column align="left" label="员工工号" prop="jobNum" width="120"/>
        <el-table-column align="left" label="合同公司" prop="company" width="240" />
        <el-table-column align="left" label="合同类型" prop="typeText" width="180" />
         <el-table-column align="left" label="合同起始日" width="150">
            <template #default="scope">{{ formatDate(scope.row.startDay) }}</template>
         </el-table-column>
         <el-table-column align="left" label="合同到期日" width="150">
            <template #default="scope">{{ formatDate(scope.row.endDay) }}</template>
         </el-table-column>
        <el-table-column align="left" label="续签次数" prop="times" width="120" />
        <el-table-column align="left" label="合同附件" prop="attachment" width="240" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateWcStaffAgreementFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
            <el-form-item label="选择员工:" prop="staffId">
              <SelectStaff v-model="formData.staffId" :disabled="type==='update'?'disabled':false">
              </SelectStaff>
            </el-form-item>
            <el-form-item label="合同公司:"  prop="company" >
              <el-input v-model="formData.company" :clearable="true"  placeholder="请输入合同公司" />
            </el-form-item>
            <el-form-item label="合同类型:"  prop="type" >
              <el-select v-model="formData.type" placeholder="选择合同类型">
                <el-option v-for="type in types" :key="type.value" :label="type.label" :value="type.value"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="合同起始日:"  prop="startDay" >
              <el-date-picker v-model="formData.startDay" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="合同到期日:"  prop="endDay" >
              <el-date-picker v-model="formData.endDay" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="续签次数:"  prop="times" >
              <el-input v-model.number="formData.times" :clearable="true" placeholder="请输入续签次数" />
            </el-form-item>
            <el-form-item label="合同附件:"  prop="attachment" >
              <el-input v-model="formData.attachment" :clearable="true"  placeholder="请输入合同附件" />
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
                <el-descriptions-item label="成员名称">
                       {{ formData.staffName }}
                </el-descriptions-item>
                <el-descriptions-item label="员工工号">
                        {{ formData.jobNum }}
                </el-descriptions-item>
                <el-descriptions-item label="合同公司">
                        {{ formData.company }}
                </el-descriptions-item>
                <el-descriptions-item label="合同类型">
                        {{ formData.typeText }}
                </el-descriptions-item>
                <el-descriptions-item label="合同起始日">
                      {{ formatDate(formData.startDay) }}
                </el-descriptions-item>
                <el-descriptions-item label="合同到期日">
                      {{ formatDate(formData.endDay) }}
                </el-descriptions-item>
                <el-descriptions-item label="续签次数">
                        {{ formData.times }}
                </el-descriptions-item>
                <el-descriptions-item label="合同附件">
                        {{ formData.attachment }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createWcStaffAgreement,
  deleteWcStaffAgreement,
  deleteWcStaffAgreementByIds,
  updateWcStaffAgreement,
  findWcStaffAgreement,
  getWcStaffAgreementList
} from '@/api/weChat/wcStaffAgreement'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import {InfoFilled, QuestionFilled} from "@element-plus/icons-vue";
import SelectStaff from "@/components/selectStaff/index.vue";

defineOptions({
    name: 'WcStaffAgreement'
})

const types = ref([
  {label: '固定期限劳动合同', value: 1},
  {label: '无固定期限劳动合同', value: 2},
  {label: '实习协议', value: 3},
  {label: '外包协议', value: 4},
  {label: '劳务派遣合同', value: 5},
  {label: '返聘协议', value: 6},
  {label: '培训协议', value: 7},
])

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        staffId: '',
        company: '',
        type: '',
        typeText: '',
        startDay: new Date(),
        endDay: new Date(),
        times: 0,
        attachment: '',
        })


// 验证规则
const rule = reactive({
               staffId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               company : [{
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
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               startDay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               endDay : [{
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
  const table = await getWcStaffAgreementList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteWcStaffAgreementFunc(row)
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
      const res = await deleteWcStaffAgreementByIds({ IDs })
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
const updateWcStaffAgreementFunc = async(row) => {
    const res = await findWcStaffAgreement({ ID: row.staffId })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rewcStaffAgreement
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWcStaffAgreementFunc = async (row) => {
    const res = await deleteWcStaffAgreement({ ID: row.ID })
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
  const res = await findWcStaffAgreement({ ID: row.staffId })
  if (res.code === 0) {
    formData.value = res.data.rewcStaffAgreement
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          staffId: '',
          company: '',
          type: '',
          typeText: '',
          startDay: new Date(),
          endDay: new Date(),
          times: 0,
          attachment: '',
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
        staffId: '',
        company: '',
        type: '',
        typeText: '',
        startDay: new Date(),
        endDay: new Date(),
        times: 0,
        attachment: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createWcStaffAgreement(formData.value)
                  break
                case 'update':
                  res = await updateWcStaffAgreement(formData.value)
                  break
                default:
                  res = await createWcStaffAgreement(formData.value)
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
