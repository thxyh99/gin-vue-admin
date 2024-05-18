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
        
        <el-table-column align="left" label="转正标题" prop="title" width="120" />
        <el-table-column align="left" label="员工ID" prop="staffId" width="120" />
         <el-table-column sortable align="left" label="入职日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.employmentDate) }}</template>
         </el-table-column>
        <el-table-column sortable align="left" label="所属部门" prop="jobDepartment" width="120" />
        <el-table-column sortable align="left" label="原职位" prop="sourcePosition" width="120" />
        <el-table-column sortable align="left" label="原职级" prop="sourceLevel" width="120" />
        <el-table-column sortable align="left" label="试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他)" prop="tryPeriod" width="120" />
        <el-table-column align="left" label="是否是部门负责人" prop="isMananger" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.isMananger) }}</template>
        </el-table-column>
        <el-table-column align="left" label="附件" prop="attachment" width="120" />
        <el-table-column align="left" label="个人自我鉴定" prop="selfOpinion" width="120" />
        <el-table-column align="left" label="用人部门意见" prop="tryOpinion" width="120" />
        <el-table-column align="left" label="转正后工作职责" prop="newResponsibilities" width="120" />
        <el-table-column sortable align="left" label="转正时间" prop="passDate" width="120" />
        <el-table-column sortable align="left" label="转正后职位" prop="newPosition" width="120" />
        <el-table-column sortable align="left" label="转正后职级" prop="newJoblevel" width="120" />
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
            <el-button type="primary" link icon="edit" class="table-button" @click="updateWcStaffPassApplicationFunc(scope.row)">变更</el-button>
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
            <el-form-item label="转正标题:"  prop="title" >
              <el-input v-model="formData.title" :clearable="true"  placeholder="请输入转正标题" />
            </el-form-item>
            <el-form-item label="员工ID:"  prop="staffId" >
              <el-input v-model.number="formData.staffId" :clearable="true" placeholder="请输入员工ID" />
            </el-form-item>
            <el-form-item label="入职日期:"  prop="employmentDate" >
              <el-date-picker v-model="formData.employmentDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="所属部门:"  prop="jobDepartment" >
                <el-select v-model="formData.jobDepartment" placeholder="请选择所属部门" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="原职位:"  prop="sourcePosition" >
                <el-select v-model="formData.sourcePosition" placeholder="请选择原职位" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="原职级:"  prop="sourceLevel" >
                <el-select v-model="formData.sourceLevel" placeholder="请选择原职级" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他):"  prop="tryPeriod" >
                <el-select v-model="formData.tryPeriod" placeholder="请选择试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他)" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="是否是部门负责人:"  prop="isMananger" >
              <el-switch v-model="formData.isMananger" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="附件:"  prop="attachment" >
              <el-input v-model="formData.attachment" :clearable="true"  placeholder="请输入附件" />
            </el-form-item>
            <el-form-item label="个人自我鉴定:"  prop="selfOpinion" >
              <el-input v-model="formData.selfOpinion" :clearable="true"  placeholder="请输入个人自我鉴定" />
            </el-form-item>
            <el-form-item label="用人部门意见:"  prop="tryOpinion" >
              <el-input v-model="formData.tryOpinion" :clearable="true"  placeholder="请输入用人部门意见" />
            </el-form-item>
            <el-form-item label="转正后工作职责:"  prop="newResponsibilities" >
                <el-select v-model="formData.newResponsibilities" placeholder="请选择转正后工作职责" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [255]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="转正时间:"  prop="passDate" >
                <el-select v-model="formData.passDate" placeholder="请选择转正时间" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="转正后职位:"  prop="newPosition" >
                <el-select v-model="formData.newPosition" placeholder="请选择转正后职位" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="转正后职级:"  prop="newJoblevel" >
                <el-select v-model="formData.newJoblevel" placeholder="请选择转正后职级" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
                </el-select>
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
                <el-descriptions-item label="转正标题">
                        {{ formData.title }}
                </el-descriptions-item>
                <el-descriptions-item label="员工ID">
                        {{ formData.staffId }}
                </el-descriptions-item>
                <el-descriptions-item label="入职日期">
                      {{ formatDate(formData.employmentDate) }}
                </el-descriptions-item>
                <el-descriptions-item label="所属部门">
                        {{ formData.jobDepartment }}
                </el-descriptions-item>
                <el-descriptions-item label="原职位">
                        {{ formData.sourcePosition }}
                </el-descriptions-item>
                <el-descriptions-item label="原职级">
                        {{ formData.sourceLevel }}
                </el-descriptions-item>
                <el-descriptions-item label="试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他)">
                        {{ formData.tryPeriod }}
                </el-descriptions-item>
                <el-descriptions-item label="是否是部门负责人">
                    {{ formatBoolean(formData.isMananger) }}
                </el-descriptions-item>
                <el-descriptions-item label="附件">
                        {{ formData.attachment }}
                </el-descriptions-item>
                <el-descriptions-item label="个人自我鉴定">
                        {{ formData.selfOpinion }}
                </el-descriptions-item>
                <el-descriptions-item label="用人部门意见">
                        {{ formData.tryOpinion }}
                </el-descriptions-item>
                <el-descriptions-item label="转正后工作职责">
                        {{ formData.newResponsibilities }}
                </el-descriptions-item>
                <el-descriptions-item label="转正时间">
                        {{ formData.passDate }}
                </el-descriptions-item>
                <el-descriptions-item label="转正后职位">
                        {{ formData.newPosition }}
                </el-descriptions-item>
                <el-descriptions-item label="转正后职级">
                        {{ formData.newJoblevel }}
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
  createWcStaffPassApplication,
  deleteWcStaffPassApplication,
  deleteWcStaffPassApplicationByIds,
  updateWcStaffPassApplication,
  findWcStaffPassApplication,
  getWcStaffPassApplicationList
} from '@/api/employee/wcStaffPassApplication'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'WcStaffPassApplication'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        title: '',
        staffId: 0,
        employmentDate: new Date(),
        isMananger: false,
        attachment: '',
        selfOpinion: '',
        tryOpinion: '',
        submitOpinion: '',
        oaId: '',
        oaStatus: false,
        })


// 验证规则
const rule = reactive({
               staffId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               employmentDate : [{
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
               sourcePosition : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               sourceLevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               tryPeriod : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               isMananger : [{
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
               selfOpinion : [{
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
               tryOpinion : [{
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
               newResponsibilities : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               passDate : [{
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
               newJoblevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
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
               oaId : [{
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
               oaStatus : [{
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
            jobDepartment: 'job_department',
            sourcePosition: 'source_position',
            sourceLevel: 'source_level',
            tryPeriod: 'try_period',
            passDate: 'pass_date',
            newPosition: 'new_position',
            newJoblevel: 'new_joblevel',
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
    if (searchInfo.value.isMananger === ""){
        searchInfo.value.isMananger=null
    }
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
  const table = await getWcStaffPassApplicationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteWcStaffPassApplicationFunc(row)
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
      const res = await deleteWcStaffPassApplicationByIds({ IDs })
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
const updateWcStaffPassApplicationFunc = async(row) => {
    const res = await findWcStaffPassApplication({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rewcStaffPassApplication
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWcStaffPassApplicationFunc = async (row) => {
    const res = await deleteWcStaffPassApplication({ ID: row.ID })
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
  const res = await findWcStaffPassApplication({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.rewcStaffPassApplication
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          title: '',
          staffId: 0,
          employmentDate: new Date(),
          isMananger: false,
          attachment: '',
          selfOpinion: '',
          tryOpinion: '',
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
        staffId: 0,
        employmentDate: new Date(),
        isMananger: false,
        attachment: '',
        selfOpinion: '',
        tryOpinion: '',
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
                  res = await createWcStaffPassApplication(formData.value)
                  break
                case 'update':
                  res = await updateWcStaffPassApplication(formData.value)
                  break
                default:
                  res = await createWcStaffPassApplication(formData.value)
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
