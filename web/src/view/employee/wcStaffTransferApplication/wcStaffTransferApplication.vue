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
        <el-table-column align="left" label="被调动人" prop="transferStaff" width="120" />
        <el-table-column sortable align="left" label="原部门" prop="jobDepartment" width="120" />
        <el-table-column sortable align="left" label="原职位" prop="jobPosition" width="120" />
        <el-table-column sortable align="left" label="调往部门" prop="toDepartment" width="120" />
        <el-table-column sortable align="left" label="调往职位" prop="toPosition" width="120" />
        <el-table-column align="left" label="调动事由" prop="transferResult" width="120" />
        <el-table-column align="left" label="原部门意见" prop="sourceResult" width="120" />
        <el-table-column align="left" label="调往部门意见" prop="toResult" width="120" />
         <el-table-column align="left" label="调入时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.toDate) }}</template>
         </el-table-column>
        <el-table-column align="left" label="考察期" prop="inspectionPerion" width="120" />
                    <el-table-column label="附件" width="200">
                        <template #default="scope">
                             <div class="file-list">
                               <el-tag v-for="file in scope.row.employementAttment" :key="file.uid">{{file.name}}</el-tag>
                             </div>
                        </template>
                    </el-table-column>
        <el-table-column align="left" label="提交意见" prop="submitOpinion" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120" />
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
              <el-input v-model.number="formData.transferStaff" :clearable="true" placeholder="请输入被调动人" />
            </el-form-item>
            <el-form-item label="原部门:"  prop="jobDepartment" >
                <el-select v-model="formData.jobDepartment" placeholder="请选择原部门" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="原职位:"  prop="jobPosition" >
                <el-select v-model="formData.jobPosition" placeholder="请选择原职位" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="调往部门:"  prop="toDepartment" >
                <el-select v-model="formData.toDepartment" placeholder="请选择调往部门" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="调往职位:"  prop="toPosition" >
                <el-select v-model="formData.toPosition" placeholder="请选择调往职位" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
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
              <el-input v-model.number="formData.inspectionPerion" :clearable="true" placeholder="请输入考察期" />
            </el-form-item>
            <el-form-item label="附件:"  prop="employementAttment" >
                <SelectFile v-model="formData.employementAttment" />
            </el-form-item>
            <el-form-item label="提交意见:"  prop="submitOpinion" >
              <el-input v-model="formData.submitOpinion" :clearable="true"  placeholder="请输入提交意见" />
            </el-form-item>
            <el-form-item label="状态:"  prop="status" >
                <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
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
                        {{ formData.jobDepartment }}
                </el-descriptions-item>
                <el-descriptions-item label="原职位">
                        {{ formData.jobPosition }}
                </el-descriptions-item>
                <el-descriptions-item label="调往部门">
                        {{ formData.toDepartment }}
                </el-descriptions-item>
                <el-descriptions-item label="调往职位">
                        {{ formData.toPosition }}
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
                        <div class="fileBtn" v-for="(item,index) in formData.employementAttment" :key="index">
                          <el-button type="primary" text bg @click="onDownloadFile(item.url)">
                            <el-icon style="margin-right: 5px"><Download /></el-icon>
                            {{ item.name }}
                          </el-button>
                        </div>
                </el-descriptions-item>
                <el-descriptions-item label="提交意见">
                        {{ formData.submitOpinion }}
                </el-descriptions-item>
                <el-descriptions-item label="状态">
                        {{ formData.status }}
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
import { getUrl } from '@/utils/image'
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'

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
        transferStaff: 0,
        transferResult: '',
        sourceResult: '',
        toResult: '',
        toDate: new Date(),
        inspectionPerion: 0,
        employementAttment: [],
        submitOpinion: '',
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
               jobDepartment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               jobPosition : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               toDepartment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               toPosition : [{
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
               status : [{
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
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            employmentDate: 'employment_date',
            transferType: 'transfer_type',
            jobDepartment: 'job_department',
            jobPosition: 'job_position',
            toDepartment: 'to_department',
            toPosition: 'to_position',
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
          transferStaff: 0,
          transferResult: '',
          sourceResult: '',
          toResult: '',
          toDate: new Date(),
          inspectionPerion: 0,
          submitOpinion: '',
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
        transferStaff: 0,
        transferResult: '',
        sourceResult: '',
        toResult: '',
        toDate: new Date(),
        inspectionPerion: 0,
        submitOpinion: '',
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

const downloadFile = (url) => {
    window.open(getUrl(url), '_blank')
}

</script>

<style>

.file-list{
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.fileBtn{
  margin-bottom: 10px;
}

.fileBtn:last-child{
  margin-bottom: 0;
}

</style>
