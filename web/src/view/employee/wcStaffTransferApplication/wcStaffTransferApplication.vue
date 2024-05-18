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
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="调动标题" prop="title" width="120" />
        <el-table-column align="left" label="申请方" prop="applyTitle" width="120" />
         <el-table-column sortable align="left" label="申请日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.employmentDate) }}</template>
         </el-table-column>
        <el-table-column sortable align="left" label="调动类型" prop="transferType" width="120" />
        <el-table-column sortable align="left" label="被调动人" prop="transferStaff" width="120" />
        <el-table-column sortable align="left" label="原部门" prop="sourceDepartment" width="120" />
        <el-table-column sortable align="left" label="原职位" prop="sourcePosition" width="120" />
        <el-table-column sortable align="left" label="调往部门" prop="newDepartment" width="120" />
        <el-table-column sortable align="left" label="调往职位" prop="newPosition" width="120" />
        <el-table-column align="left" label="调动事由" prop="transferResult" width="120" />
        <el-table-column align="left" label="原部门意见" prop="sourceResult" width="120" />
        <el-table-column align="left" label="调往部门意见" prop="toResult" width="120" />
         <el-table-column sortable align="left" label="调入时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.toDate) }}</template>
         </el-table-column>
        <el-table-column sortable align="left" label="考察期" prop="inspectionPerion" width="120" />
        <el-table-column align="left" label="附件" prop="attachment" width="120" />
        <el-table-column align="left" label="提交意见" prop="submitOpinion" width="120" />
        <el-table-column align="left" label="OAID" prop="oaId" width="120" />
        <el-table-column align="left" label="OA状态" prop="oaStatus" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.oaStatus) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateWcStaffTransferApplicationFunc(scope.row)">变更</el-button>
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
            <el-form-item label="调动标题:"  prop="title" >
              <el-input v-model="formData.title" :clearable="true"  placeholder="请输入调动标题" />
            </el-form-item>
            <el-form-item label="申请方:"  prop="applyTitle" >
              <el-input v-model="formData.applyTitle" :clearable="true"  placeholder="请输入申请方" />
            </el-form-item>
            <el-form-item label="申请日期:"  prop="employmentDate" >
              <el-date-picker v-model="formData.employmentDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="调动类型:"  prop="transferType" >
                <el-select v-model="formData.transferType" placeholder="请选择调动类型" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="被调动人:"  prop="transferStaff" >
                <el-select v-model="formData.transferStaff" placeholder="请选择被调动人" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="原部门:"  prop="sourceDepartment" >
                <el-select v-model="formData.sourceDepartment" placeholder="请选择原部门" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="原职位:"  prop="sourcePosition" >
                <el-select v-model="formData.sourcePosition" placeholder="请选择原职位" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="调往部门:"  prop="newDepartment" >
                <el-select v-model="formData.newDepartment" placeholder="请选择调往部门" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="调往职位:"  prop="newPosition" >
                <el-select v-model="formData.newPosition" placeholder="请选择调往职位" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="调动事由:"  prop="transferResult" >
              <el-input v-model="formData.transferResult" :clearable="true"  placeholder="请输入调动事由" />
            </el-form-item>
            <el-form-item label="原部门意见:"  prop="sourceResult" >
              <el-input v-model="formData.sourceResult" :clearable="true"  placeholder="请输入原部门意见" />
            </el-form-item>
            <el-form-item label="调往部门意见:"  prop="toResult" >
              <el-input v-model="formData.toResult" :clearable="true"  placeholder="请输入调往部门意见" />
            </el-form-item>
            <el-form-item label="调入时间:"  prop="toDate" >
              <el-date-picker v-model="formData.toDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="考察期:"  prop="inspectionPerion" >
                <el-select v-model="formData.inspectionPerion" placeholder="请选择考察期" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="附件:"  prop="attachment" >
              <el-input v-model="formData.attachment" :clearable="true"  placeholder="请输入附件" />
            </el-form-item>
            <el-form-item label="提交意见:"  prop="submitOpinion" >
              <el-input v-model="formData.submitOpinion" :clearable="true"  placeholder="请输入提交意见" />
            </el-form-item>
            <el-form-item label="OAID:"  prop="oaId" >
              <el-input v-model="formData.oaId" :clearable="true"  placeholder="请输入OAID" />
            </el-form-item>
            <el-form-item label="OA状态:"  prop="oaStatus" >
              <el-switch v-model="formData.oaStatus" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
                <el-descriptions-item label="调动标题">
                        {{ formData.title }}
                </el-descriptions-item>
                <el-descriptions-item label="申请方">
                        {{ formData.applyTitle }}
                </el-descriptions-item>
                <el-descriptions-item label="申请日期">
                      {{ formatDate(formData.employmentDate) }}
                </el-descriptions-item>
                <el-descriptions-item label="调动类型">
                        {{ formData.transferType }}
                </el-descriptions-item>
                <el-descriptions-item label="被调动人">
                        {{ formData.transferStaff }}
                </el-descriptions-item>
                <el-descriptions-item label="原部门">
                        {{ formData.sourceDepartment }}
                </el-descriptions-item>
                <el-descriptions-item label="原职位">
                        {{ formData.sourcePosition }}
                </el-descriptions-item>
                <el-descriptions-item label="调往部门">
                        {{ formData.newDepartment }}
                </el-descriptions-item>
                <el-descriptions-item label="调往职位">
                        {{ formData.newPosition }}
                </el-descriptions-item>
                <el-descriptions-item label="调动事由">
                        {{ formData.transferResult }}
                </el-descriptions-item>
                <el-descriptions-item label="原部门意见">
                        {{ formData.sourceResult }}
                </el-descriptions-item>
                <el-descriptions-item label="调往部门意见">
                        {{ formData.toResult }}
                </el-descriptions-item>
                <el-descriptions-item label="调入时间">
                      {{ formatDate(formData.toDate) }}
                </el-descriptions-item>
                <el-descriptions-item label="考察期">
                        {{ formData.inspectionPerion }}
                </el-descriptions-item>
                <el-descriptions-item label="附件">
                        {{ formData.attachment }}
                </el-descriptions-item>
                <el-descriptions-item label="提交意见">
                        {{ formData.submitOpinion }}
                </el-descriptions-item>
                <el-descriptions-item label="OAID">
                        {{ formData.oaId }}
                </el-descriptions-item>
                <el-descriptions-item label="OA状态">
                    {{ formatBoolean(formData.oaStatus) }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createWcStaffTransferApplication,
  deleteWcStaffTransferApplication,
  deleteWcStaffTransferApplicationByIds,
  updateWcStaffTransferApplication,
  findWcStaffTransferApplication,
  getWcStaffTransferApplicationList
} from '@/api/employee/wcStaffTransferApplication'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'WcStaffTransferApplication'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        title: '',
        applyTitle: '',
        employmentDate: new Date(),
        transferResult: '',
        sourceResult: '',
        toResult: '',
        toDate: new Date(),
        attachment: '',
        submitOpinion: '',
        oaId: '',
        oaStatus: false,
        })


// 验证规则
const rule = reactive({
               applyTitle : [{
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
               employmentDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               transferType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               transferStaff : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               sourceDepartment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               sourcePosition : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               newDepartment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               newPosition : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               transferResult : [{
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
               sourceResult : [{
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
               toResult : [{
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
               toDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               inspectionPerion : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               attachment : [{
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
               submitOpinion : [{
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
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            employmentDate: 'employment_date',
            transferType: 'transfer_type',
            transferStaff: 'transfer_staff',
            sourceDepartment: 'source_department',
            sourcePosition: 'source_position',
            newDepartment: 'new_department',
            newPosition: 'new_position',
            toDate: 'to_date',
            inspectionPerion: 'inspection_perion',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}

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
    if (searchInfo.value.oaStatus === ""){
        searchInfo.value.oaStatus=null
    }
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
  const table = await getWcStaffTransferApplicationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteWcStaffTransferApplicationFunc(row)
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
      const res = await deleteWcStaffTransferApplicationByIds({ IDs })
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
const updateWcStaffTransferApplicationFunc = async(row) => {
    const res = await findWcStaffTransferApplication({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rewcStaffTransferApplication
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWcStaffTransferApplicationFunc = async (row) => {
    const res = await deleteWcStaffTransferApplication({ ID: row.ID })
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
  const res = await findWcStaffTransferApplication({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.rewcStaffTransferApplication
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          title: '',
          applyTitle: '',
          employmentDate: new Date(),
          transferResult: '',
          sourceResult: '',
          toResult: '',
          toDate: new Date(),
          attachment: '',
          submitOpinion: '',
          oaId: '',
          oaStatus: false,
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
        title: '',
        applyTitle: '',
        employmentDate: new Date(),
        transferResult: '',
        sourceResult: '',
        toResult: '',
        toDate: new Date(),
        attachment: '',
        submitOpinion: '',
        oaId: '',
        oaStatus: false,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createWcStaffTransferApplication(formData.value)
                  break
                case 'update':
                  res = await updateWcStaffTransferApplication(formData.value)
                  break
                default:
                  res = await createWcStaffTransferApplication(formData.value)
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
