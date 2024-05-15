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
        <el-table-column align="left" label="员工类型" prop="typeText" width="120" />
        <el-table-column align="left" label="员工状态" prop="statusText" width="120" />
        <el-table-column align="left" label="部门信息" prop="department" width="200"/>
        <el-table-column align="left" label="职务信息" prop="position" width="200"/>
        <el-table-column align="left" label="职级类型" prop="rankTypeText" width="120" />
        <el-table-column align="left" label="职级" prop="rankText" width="120" />
        <el-table-column align="left" label="等级工资" prop="rankSalary" width="120" />
        <el-table-column align="left" label="费用科目" prop="expenseAccountText" width="120" />
        <el-table-column align="left" label="入职日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.employmentDate) }}</template>
        </el-table-column>
        <el-table-column align="left" label="试用期" prop="tryPeriodText" width="120" />
        <el-table-column align="left" label="转正日期" width="180">
         <template #default="scope">{{ formatDate(scope.row.formalDate) }}</template>
         </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateWcStaffJobFunc(scope.row)">变更</el-button>
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
            <el-form-item label="员工类型:"  prop="type" >
              <el-select v-model="formData.type" placeholder="选择员工类型">
                <el-option v-for="type in types" :key="type.value" :label="type.label" :value="type.value"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="员工状态:"  prop="status" >
              <el-select v-model="formData.status" placeholder="选择员工状态">
                <el-option v-for="status in statusList" :key="status.value" :label="status.label" :value="status.value"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="部门信息:" prop="departmentIds">
              <SelectDepartment v-model="formData.departmentIds">
              </SelectDepartment>
            </el-form-item>
            <el-form-item label="职务信息:" prop="positionIds">
              <SelectPosition v-model="formData.positionIds">
              </SelectPosition>
            </el-form-item>
            <el-form-item label="职级类型:"  prop="rankType" >
              <SelectRankType v-model="formData.rankType" placeholder="选择职级类型" @change="handleTypeChange">
              </SelectRankType>
            </el-form-item>
            <el-form-item label="职级:"  prop="rank" >
              <SelectRank ref="SelectRankRef" v-model="formData.rank" placeholder="选择职级">
              </SelectRank>
            </el-form-item>
            <el-form-item label="等级工资:"  prop="rankSalary" >
              <el-input-number v-model="formData.rankSalary"  style="width:100%" :precision="0" :clearable="true"  />
            </el-form-item>
            <el-form-item label="费用科目:"  prop="expenseAccount" >
              <el-select v-model="formData.expenseAccount" placeholder="选择费用科目">
                <el-option v-for="expenseAccount in expenseAccounts" :key="expenseAccount.value" :label="expenseAccount.label" :value="expenseAccount.value"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="入职日期:"  prop="employmentDate" >
              <el-date-picker v-model="formData.employmentDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="试用期:"  prop="tryPeriod" >
              <el-select v-model="formData.tryPeriod" placeholder="选择试用期">
                <el-option v-for="tryPeriod in tryPeriods" :key="tryPeriod.value" :label="tryPeriod.label" :value="tryPeriod.value"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="转正日期:"  prop="formalDate" >
              <el-date-picker v-model="formData.formalDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
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
                <el-descriptions-item label="员工类型">
                        {{ formData.typeText }}
                </el-descriptions-item>
                <el-descriptions-item label="员工状态">
                        {{ formData.statusText }}
                </el-descriptions-item>
                <el-descriptions-item label="部门信息">
                        {{ formData.department }}
                </el-descriptions-item>
                <el-descriptions-item label="职务信息">
                        {{ formData.position }}
                </el-descriptions-item>
                <el-descriptions-item label="职级类型">
                        {{ formData.rankTypeText }}
                </el-descriptions-item>
                <el-descriptions-item label="职级">
                        {{ formData.rankText }}
                </el-descriptions-item>
                <el-descriptions-item label="等级工资">
                        {{ formData.rankSalary }}
                </el-descriptions-item>
                <el-descriptions-item label="费用科目">
                        {{ formData.expenseAccountText }}
                </el-descriptions-item>
                <el-descriptions-item label="试用期">
                        {{ formData.tryPeriodText }}
                </el-descriptions-item>
                <el-descriptions-item label="入职日期">
                      {{ formatDate(formData.employmentDate) }}
                </el-descriptions-item>
                <el-descriptions-item label="转正日期">
                      {{ formatDate(formData.formalDate) }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createWcStaffJob,
  deleteWcStaffJob,
  deleteWcStaffJobByIds,
  updateWcStaffJob,
  findWcStaffJob,
  getWcStaffJobList
} from '@/api/weChat/wcStaffJob'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import {InfoFilled, QuestionFilled} from "@element-plus/icons-vue";
import SelectStaff from "@/components/selectStaff/index.vue";
import SelectPosition from "@/components/selectPosition/index.vue";
import SelectDepartment from "@/components/selectDepartment/index.vue";
import SelectRankType from "@/components/selectRankType/index.vue";
import SelectRank from "@/components/selectRank/index.vue";

defineOptions({
    name: 'WcStaffJob'
})

const types = ref([
  { label: '其他', value: 0 },
  { label: '全职', value: 1 },
  { label: '兼职', value: 2 },
  { label: '跟岗实习', value: 3 },
  { label: '顶岗实习', value: 4 },
  { label: '退休返聘', value: 5 },
  { label: '劳务外包', value: 6 },
])

const statusList = ref([
  { label: '试用', value: 1 },
  { label: '正式', value: 2 },
  { label: '待离职', value: 3 },
  { label: '已离职', value: 4 },
])

const tryPeriods = ref([
  { label: '其他', value: 0 },
  { label: '无试用期', value: 1 },
  { label: '2个月', value: 2 },
  { label: '6个月', value: 3 },
])

const expenseAccounts = ref([
  { label: '管理费用', value: 1 },
  { label: '研发费用', value: 2 },
  { label: '生产费用', value: 3 },
  { label: '销售费用', value: 4 },
])

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        staffId: '',
        type: '',
        status: '',
        tryPeriod: '',
        employmentDate: new Date(),
        formalDate: new Date(),
        expenseAccount: '',
        staffName:'',
        jobNum: '',
        typeText:'',
        statusText:'',
        tryPeriodText:'',
        expenseAccountText:'',
        rankType:'',
        rankTypeText:'',
        rank:'',
        rankText:'',
        rankSalary:0,
        department:'',
        departmentIds:[],
        position:'',
        positionIds:[],
        })


// 验证规则
const rule = reactive({
                staffId : [{
                  required: true,
                  message: '',
                  trigger: ['input','blur'],
                },
                ],
               type : [{
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
               tryPeriod : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
                expenseAccount : [{
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
               formalDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
              rankType : [{
                required: true,
                message: '',
                trigger: ['input','blur'],
              },
              ],
              rank : [{
                required: true,
                message: '',
                trigger: ['input','blur'],
              },
              ],
              departmentIds : [{
                required: true,
                message: '',
                trigger: ['input','blur'],
              },
              ],
              positionIds : [{
                required: true,
                message: '',
                trigger: ['input','blur'],
              },
              ],
              rankSalary : [{
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
const rankList = ref({})

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

const SelectRankRef = ref()
const handleTypeChange = (e) => {
  console.log(e)
  console.log(SelectRankRef.value)
  SelectRankRef.value.initSelectOptions(e)
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
  const table = await getWcStaffJobList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteWcStaffJobFunc(row)
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
      const res = await deleteWcStaffJobByIds({ IDs })
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
const updateWcStaffJobFunc = async(row) => {
    const res = await findWcStaffJob({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rewcStaffJob
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWcStaffJobFunc = async (row) => {
    const res = await deleteWcStaffJob({ ID: row.ID })
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
  const res = await findWcStaffJob({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.rewcStaffJob
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
            staffId: '',
            type: '',
            status: '',
            tryPeriod: '',
            employmentDate: new Date(),
            formalDate: new Date(),
            expenseAccount: '',
            staffName:'',
            jobNum: '',
            typeText:'',
            statusText:'',
            tryPeriodText:'',
            expenseAccountText:'',
            rankType:'',
            rankTypeText:'',
            rank:'',
            rankText:'',
            rankSalary:0,
            department:'',
            departmentIds:[],
            position:'',
            positionIds:[],
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
          type: '',
          status: '',
          tryPeriod: '',
          employmentDate: new Date(),
          formalDate: new Date(),
          expenseAccount: '',
          staffName:'',
          jobNum: '',
          typeText:'',
          statusText:'',
          tryPeriodText:'',
          expenseAccountText:'',
          rankType:'',
          rankTypeText:'',
          rank:'',
          rankText:'',
          rankSalary:0,
          department:'',
          departmentIds:[],
          position:'',
          positionIds:[],
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createWcStaffJob(formData.value)
                  break
                case 'update':
                  res = await updateWcStaffJob(formData.value)
                  break
                default:
                  res = await createWcStaffJob(formData.value)
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
