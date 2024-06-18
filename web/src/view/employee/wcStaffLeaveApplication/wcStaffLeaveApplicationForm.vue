
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="离职标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入离职标题" />
       </el-form-item>
        <el-form-item label="员工ID:" prop="staffId">
          <el-input v-model.number="formData.staffId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="解除日期:" prop="leaveDate">
          <el-date-picker v-model="formData.leaveDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="所属部门:" prop="jobDepartment">
          <el-input v-model.number="formData.jobDepartment" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="离职类型:" prop="leaveType">
          <el-switch v-model="formData.leaveType" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="事由:" prop="leaveResult">
          <el-input v-model="formData.leaveResult" :clearable="true"  placeholder="请输入事由" />
       </el-form-item>
        <el-form-item label="申请表:" prop="attachment">
          <el-input v-model="formData.attachment" :clearable="true"  placeholder="请输入申请表" />
       </el-form-item>
        <el-form-item label="交接清单:" prop="checkList">
          <el-input v-model="formData.checkList" :clearable="true"  placeholder="请输入交接清单" />
       </el-form-item>
        <el-form-item label="是否开具离职证明:" prop="isLeave">
          <el-switch v-model="formData.isLeave" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="是否入住公司宿舍:" prop="isHavedorm">
          <el-switch v-model="formData.isHavedorm" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="宿舍所在地:" prop="dormLocation">
          <el-input v-model="formData.dormLocation" :clearable="true"  placeholder="请输入宿舍所在地" />
       </el-form-item>
        <el-form-item label="房间门牌号:" prop="roomNum">
          <el-input v-model="formData.roomNum" :clearable="true"  placeholder="请输入房间门牌号" />
       </el-form-item>
        <el-form-item label="提交意见:" prop="submitOpinion">
          <el-input v-model="formData.submitOpinion" :clearable="true"  placeholder="请输入提交意见" />
       </el-form-item>
        <el-form-item label="OAID:" prop="oaId">
          <el-input v-model="formData.oaId" :clearable="true"  placeholder="请输入OAID" />
       </el-form-item>
        <el-form-item label="OA状态:" prop="oaStatus">
          <el-input v-model.number="formData.oaStatus" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createWcStaffLeaveApplication,
  updateWcStaffLeaveApplication,
  findWcStaffLeaveApplication
} from '@/api/employee/wcStaffLeaveApplication'

defineOptions({
    name: 'WcStaffLeaveApplicationForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            title: '',
            staffId: 0,
            leaveDate: new Date(),
            jobDepartment: 0,
            leaveType: false,
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
// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               staffId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               leaveDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               jobDepartment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               leaveType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               leaveResult : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWcStaffLeaveApplication({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcStaffLeaveApplication
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
