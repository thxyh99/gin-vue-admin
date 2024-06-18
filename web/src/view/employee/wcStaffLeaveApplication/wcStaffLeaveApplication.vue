
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
        
        <el-table-column align="left" label="离职标题" prop="title" width="120" />
        <el-table-column align="left" label="员工ID" prop="staffId" width="120" />
         <el-table-column sortable align="left" label="解除日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.leaveDate) }}</template>
         </el-table-column>
        <el-table-column sortable align="left" label="所属部门" prop="jobDepartment" width="120" />
        <el-table-column sortable align="left" label="职务" prop="jobPosition" width="120" />
        <el-table-column sortable align="left" label="离职类型" prop="leaveType" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.leaveType) }}</template>
        </el-table-column>
        <el-table-column align="left" label="事由" prop="leaveResult" width="120" />
        <el-table-column align="left" label="申请表" prop="attachment" width="120" />
        <el-table-column align="left" label="交接清单" prop="checkList" width="120" />
        <el-table-column align="left" label="是否开具离职证明" prop="isLeave" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.isLeave) }}</template>
        </el-table-column>
        <el-table-column align="left" label="是否入住公司宿舍" prop="isHavedorm" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.isHavedorm) }}</template>
        </el-table-column>
        <el-table-column align="left" label="宿舍所在地" prop="dormLocation" width="120" />
        <el-table-column align="left" label="房间门牌号" prop="roomNum" width="120" />
        <el-table-column align="left" label="提交意见" prop="submitOpinion" width="120" />
         <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateWcStaffLeaveApplicationFunc(scope.row)">变更</el-button>
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
            <el-form-item label="员工ID:"  prop="staffId" >
              <SelectStaff v-model="formData.staffId" :disabled="type==='update'?'disabled':false" @changeSuccess="changeSuccess">
              </SelectStaff>
            </el-form-item>
            <el-form-item label="离职标题:"  prop="title" >
              <el-input v-model="formData.title" :clearable="true"  placeholder="请输入离职标题" />
            </el-form-item>
            <el-form-item label="解除日期:"  prop="leaveDate" >
              <el-date-picker v-model="formData.leaveDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="所属部门:"  prop="jobDepartment" >
              <SelectDepartment v-model="formData.jobDepartment">
              </SelectDepartment>
            </el-form-item>
            <el-form-item label="职务:"  prop="jobPosition" >
              <SelectPosition v-model="formData.jobPosition">
              </SelectPosition>
            </el-form-item>
            <el-form-item label="离职类型:"  prop="leaveType" >
              <el-select v-model="formData.leaveType" placeholder="请选择离职类型" clearable>
                <el-option v-for="item in leaveTypes" :key="item.value" :label="item.label" :value="item.value"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="事由:"  prop="leaveResult" >
              <el-input v-model="formData.leaveResult" :clearable="true"  placeholder="请输入事由" />
            </el-form-item>
            <el-form-item label="申请表:"  prop="attachment" >
              <el-input v-model="formData.attachment" :clearable="true"  placeholder="请输入申请表" />
            </el-form-item>
            <el-form-item label="交接清单:"  prop="checkList" >
              <el-input v-model="formData.checkList" :clearable="true"  placeholder="请输入交接清单" />
            </el-form-item>
            <el-form-item label="是否开具离职证明:"  prop="isLeave" >
              <el-switch v-model="formData.isLeave" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="是否入住公司宿舍:"  prop="isHavedorm" >
              <el-switch v-model="formData.isHavedorm" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="宿舍所在地:"  prop="dormLocation" >
              <el-input v-model="formData.dormLocation" :clearable="true"  placeholder="请输入宿舍所在地" />
            </el-form-item>
            <el-form-item label="房间门牌号:"  prop="roomNum" >
              <el-input v-model="formData.roomNum" :clearable="true"  placeholder="请输入房间门牌号" />
            </el-form-item>
            <el-form-item label="提交意见:"  prop="submitOpinion" >
              <el-input v-model="formData.submitOpinion" :clearable="true"  placeholder="请输入提交意见" />
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
                <el-descriptions-item label="离职标题">
                        {{ formData.title }}
                </el-descriptions-item>
                <el-descriptions-item label="员工ID">
                        {{ formData.staffId }}
                </el-descriptions-item>
                <el-descriptions-item label="解除日期">
                      {{ formatDate(formData.leaveDate) }}
                </el-descriptions-item>
                <el-descriptions-item label="所属部门">
                        {{ formData.jobDepartment }}
                </el-descriptions-item>
                <el-descriptions-item label="职务">
                        {{ formData.jobPosition }}
                </el-descriptions-item>
                <el-descriptions-item label="离职类型">
                    {{ formatData.leaveType }}
                </el-descriptions-item>
                <el-descriptions-item label="事由">
                        {{ formData.leaveResult }}
                </el-descriptions-item>
                <el-descriptions-item label="申请表">
                        {{ formData.attachment }}
                </el-descriptions-item>
                <el-descriptions-item label="交接清单">
                        {{ formData.checkList }}
                </el-descriptions-item>
                <el-descriptions-item label="是否开具离职证明">
                    {{ formatBoolean(formData.isLeave) }}
                </el-descriptions-item>
                <el-descriptions-item label="是否入住公司宿舍">
                    {{ formatBoolean(formData.isHavedorm) }}
                </el-descriptions-item>
                <el-descriptions-item label="宿舍所在地">
                        {{ formData.dormLocation }}
                </el-descriptions-item>
                <el-descriptions-item label="房间门牌号">
                        {{ formData.roomNum }}
                </el-descriptions-item>
                <el-descriptions-item label="提交意见">
                        {{ formData.submitOpinion }}
                </el-descriptions-item>
                <el-descriptions-item label="OAID">
                        {{ formData.oaId }}
                </el-descriptions-item>
                <el-descriptions-item label="OA状态">
                        {{ formData.oaStatus }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createWcStaffLeaveApplication,
  deleteWcStaffLeaveApplication,
  deleteWcStaffLeaveApplicationByIds,
  updateWcStaffLeaveApplication,
  findWcStaffLeaveApplication,
  getWcStaffLeaveApplicationList
} from '@/api/employee/wcStaffLeaveApplication'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

import SelectStaff from "@/components/selectStaff/index.vue";
import SelectDepartment from "@/components/selectDepartment/index.vue";
import SelectPosition from "@/components/selectPosition/index.vue";

defineOptions({
    name: 'WcStaffLeaveApplication'
})

const leaveTypes = ref([
  { label: '自离', value: '自离' },
  { label: '辞退', value: '辞退' },
  { label: '协商解除', value: '协商解除' },
  { label: '退休', value: '退休' },
  { label: '其他', value: '其他' },
])


// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        title: '',
        staffId: 0,
        leaveDate: new Date(),
        jobDepartment: 0,
        jobPosition: 0,
        leaveType: '',
        leaveResult: '',
        attachment: '',
        checkList: '',
        isLeave: false,
        isHavedorm: false,
        dormLocation: '',
        roomNum: '',
        submitOpinion: '',
        oaId: '',
        oaStatus: 0,
        })

const changeSuccess = e => {
  formData.value.title = '离职申请-' + e.label
}

// 验证规则
const rule = reactive({
               title : [{
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
               staffId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               leaveDate : [{
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
               leaveType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               leaveResult : [{
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
            leaveDate: 'leave_date',
            jobDepartment: 'job_department',
            leaveType: 'leave_type',
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
    if (searchInfo.value.leaveType === ""){
        searchInfo.value.leaveType=null
    }
    if (searchInfo.value.isLeave === ""){
        searchInfo.value.isLeave=null
    }
    if (searchInfo.value.isHavedorm === ""){
        searchInfo.value.isHavedorm=null
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
  const table = await getWcStaffLeaveApplicationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteWcStaffLeaveApplicationFunc(row)
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
      const res = await deleteWcStaffLeaveApplicationByIds({ IDs })
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
const updateWcStaffLeaveApplicationFunc = async(row) => {
    const res = await findWcStaffLeaveApplication({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rewcStaffLeaveApplication
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWcStaffLeaveApplicationFunc = async (row) => {
    const res = await deleteWcStaffLeaveApplication({ ID: row.ID })
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
  const res = await findWcStaffLeaveApplication({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.rewcStaffLeaveApplication
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          title: '',
          staffId: 0,
          leaveDate: new Date(),
          jobDepartment: 0,
          jobPosition: 0,
          leaveType: '',
          leaveResult: '',
          attachment: '',
          checkList: '',
          isLeave: false,
          isHavedorm: false,
          dormLocation: '',
          roomNum: '',
          submitOpinion: '',
          oaId: '',
          oaStatus: 0,
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
        staffId: 0,
        leaveDate: new Date(),
        jobDepartment: 0,
        jobPosition: 0,
        leaveType: '',
        leaveResult: '',
        attachment: '',
        checkList: '',
        isLeave: false,
        isHavedorm: false,
        dormLocation: '',
        roomNum: '',
        submitOpinion: '',
        oaId: '',
        oaStatus: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createWcStaffLeaveApplication(formData.value)
                  break
                case 'update':
                  res = await updateWcStaffLeaveApplication(formData.value)
                  break
                default:
                  res = await createWcStaffLeaveApplication(formData.value)
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

