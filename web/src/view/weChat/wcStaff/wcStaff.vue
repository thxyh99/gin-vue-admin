<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
               @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled/></el-icon>
          </el-tooltip>
        </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期"
                          :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期"
                          :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
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
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">
          删除
        </el-button>
        <ExportTemplate
            template-id="staff"
        />
        <ExportExcel
            template-id="staff"
            :limit="9999"
        />
        <ImportExcel
            template-id="staff"
            @on-success="getTableData"
        />
      </div>
      <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
          @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55"/>

        <!--        <el-table-column align="left" label="日期" width="180">-->
        <!--            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>-->
        <!--        </el-table-column>-->

        <el-table-column align="left" label="成员名称" prop="name" width="180"/>
        <el-table-column align="left" label="员工工号" prop="jobNum" width="150"/>
        <el-table-column align="left" label="企微账号" prop="userid" width="150"/>
        <el-table-column align="left" label="部门信息" prop="department" width="200"/>
        <el-table-column align="left" label="职务信息" prop="position" width="200"/>
        <el-table-column align="left" label="性别" prop="genderText" width="90"/>
        <el-table-column align="left" label="是否领导" prop="isLeaderText" width="90"/>
        <el-table-column align="left" label="手机" prop="mobile" width="150"/>
<!--        <el-table-column align="left" label="座机" prop="telephone" width="120"/>-->
<!--        <el-table-column align="left" label="个人邮箱" prop="email" width="130"/>-->
<!--        <el-table-column align="left" label="地址" prop="address" width="120"/>-->
<!--        <el-table-column align="left" label="企业邮箱" prop="bizMail" width="120"/>-->
        <el-table-column align="left" label="状态" prop="statusText" width="120"/>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled/>
              </el-icon>
              查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateWcStaffFunc(scope.row)">变更
            </el-button>
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
          <span class="text-lg">{{ type === 'create' ? '添加' : '修改' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="成员名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入成员名称"/>
        </el-form-item>
        <el-form-item label="员工工号:" prop="jobNum">
          <el-input v-model="formData.jobNum" :clearable="true" placeholder="请输入员工工号"/>
        </el-form-item>
        <el-form-item label="企微账号:" prop="userid">
          <el-input v-model="formData.userid" :clearable="true" placeholder="请输入企微账号"/>
        </el-form-item>
        <el-form-item label="职务信息:" prop="positionIds">
          <SelectPosition v-model="formData.positionIds">
          </SelectPosition>
        </el-form-item>
        <el-form-item label="性别:" prop="gender">
          <el-radio-group  v-model="formData.gender">
            <el-radio :label=1>男</el-radio>
            <el-radio :label=2>女</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="是否领导:" prop="isLeader">
          <el-radio-group v-model="formData.isLeader">
            <el-radio :label=1>是</el-radio>
            <el-radio :label=0>否</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="手机:" prop="mobile">
          <el-input v-model="formData.mobile" :clearable="true" placeholder="请输入手机"/>
        </el-form-item>
        <el-form-item label="座机:" prop="telephone">
          <el-input v-model="formData.telephone" :clearable="true" placeholder="请输入座机"/>
        </el-form-item>
        <el-form-item label="个人邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true" placeholder="请输入个人邮箱"/>
        </el-form-item>
        <el-form-item label="地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true" placeholder="请输入地址"/>
        </el-form-item>
        <el-form-item label="企业邮箱:" prop="bizMail">
          <el-input v-model="formData.bizMail" :clearable="true" placeholder="请输入企业邮箱"/>
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择状态">
            <el-option v-for="status in statuses" :key="status.value" :label="status.label" :value="status.value"></el-option>
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
        <el-descriptions-item label="成员名称">
          {{ formData.name }}
        </el-descriptions-item>
        <el-descriptions-item label="员工工号">
          {{ formData.jobNum }}
        </el-descriptions-item>
        <el-descriptions-item label="企微账号">
          {{ formData.userid }}
        </el-descriptions-item>
        <el-descriptions-item label="部门信息">
          {{ formData.department }}
        </el-descriptions-item>
        <el-descriptions-item label="职务信息">
          {{ formData.position }}
        </el-descriptions-item>
        <el-descriptions-item label="性别">
          {{ formData.genderText }}
        </el-descriptions-item>
        <el-descriptions-item label="是否领导">
          {{ formData.isLeaderText }}
        </el-descriptions-item>
        <el-descriptions-item label="手机">
          {{ formData.mobile }}
        </el-descriptions-item>
        <el-descriptions-item label="座机">
          {{ formData.telephone }}
        </el-descriptions-item>
        <el-descriptions-item label="个人邮箱">
          {{ formData.email }}
        </el-descriptions-item>
        <el-descriptions-item label="地址">
          {{ formData.address }}
        </el-descriptions-item>
        <el-descriptions-item label="企业邮箱">
          {{ formData.bizMail }}
        </el-descriptions-item>
        <el-descriptions-item label="状态">
          {{ formData.statusText }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createWcStaff,
  deleteWcStaff,
  deleteWcStaffByIds,
  updateWcStaff,
  findWcStaff,
  getWcStaffList
} from '@/api/weChat/wcStaff'

// 全量引入格式化工具 请按需保留
import {getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile} from '@/utils/format'
import {ElMessage, ElMessageBox} from 'element-plus'
import {ref, reactive} from 'vue'
import ImportExcel from "@/components/exportExcel/wechat/importExcel.vue";
import ExportExcel from "@/components/exportExcel/wechat/exportExcel.vue";
import ExportTemplate from "@/components/exportExcel/wechat/exportTemplate.vue";
import SelectPosition from "@/components/selectPosition/index.vue";

defineOptions({
  name: 'WcStaff'
})

const statuses = ref([
  { label: '已激活', value: 1 },
  { label: '已禁用', value: 2 },
  { label: '未激活', value: 4 },
  { label: '退出企业', value: 5 }
])

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  userId: 0,
  userid: '',
  jobNum: '',
  name: '',
  department: '',
  positionIds: [],
  gender: 0,
  isLeader: 0,
  mobile: '',
  telephone: '',
  email: '',
  address: '',
  bizMail: '',
  status: 1,
})


// 验证规则
const rule = reactive({
  userId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  userid: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
    {
      whitespace: true,
      message: '不能只输入空格',
      trigger: ['input', 'blur'],
    }
  ],
  jobNum: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
    {
      whitespace: true,
      message: '不能只输入空格',
      trigger: ['input', 'blur'],
    }
  ],
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
    {
      whitespace: true,
      message: '不能只输入空格',
      trigger: ['input', 'blur'],
    }
  ],
  gender: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  isLeader: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  mobile: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
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
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change'
    }
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
  elSearchFormRef.value?.validate(async (valid) => {
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
const getTableData = async () => {
  const table = await getWcStaffList({page: page.value, pageSize: pageSize.value, ...searchInfo.value})
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
const setOptions = async () => {
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
    deleteWcStaffFunc(row)
  })
}

// 多选删除
const onDelete = async () => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
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
    const res = await deleteWcStaffByIds({IDs})
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
const updateWcStaffFunc = async (row) => {
  const res = await findWcStaff({ID: row.ID})
  type.value = 'update'
  if (res.code === 0) {
    // console.log(formData.value)
    formData.value = res.data.rewcStaff
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteWcStaffFunc = async (row) => {
  const res = await deleteWcStaff({ID: row.ID})
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
  const res = await findWcStaff({ID: row.ID})
  if (res.code === 0) {
    formData.value = res.data.rewcStaff
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    userId: 0,
    userid: '',
    jobNum: '',
    name: '',
    department: '',
    position: '',
    gender: 0,
    isLeader: 0,
    mobile: '',
    telephone: '',
    email: '',
    address: '',
    bizMail: '',
    status: 1,
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
    userId: 0,
    userid: '',
    name: '',
    department: '',
    position: '',
    gender: 0,
    isLeader: 0,
    mobile: '',
    telephone: '',
    email: '',
    address: '',
    bizMail: '',
    status: 1,
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createWcStaff(formData.value)
        break
      case 'update':
        res = await updateWcStaff(formData.value)
        break
      default:
        res = await createWcStaff(formData.value)
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
