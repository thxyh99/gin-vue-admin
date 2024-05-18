<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="调动标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入调动标题" />
       </el-form-item>
        <el-form-item label="申请方:" prop="applyTitle">
          <el-input v-model="formData.applyTitle" :clearable="true"  placeholder="请输入申请方" />
       </el-form-item>
        <el-form-item label="申请日期:" prop="employmentDate">
          <el-date-picker v-model="formData.employmentDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="调动类型:" prop="transferType">
        <el-select v-model="formData.transferType" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="被调动人:" prop="transferStaff">
        <el-select v-model="formData.transferStaff" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="原部门:" prop="sourceDepartment">
        <el-select v-model="formData.sourceDepartment" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="原职位:" prop="sourcePosition">
        <el-select v-model="formData.sourcePosition" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="调往部门:" prop="newDepartment">
        <el-select v-model="formData.newDepartment" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="调往职位:" prop="newPosition">
        <el-select v-model="formData.newPosition" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="调动事由:" prop="transferResult">
          <el-input v-model="formData.transferResult" :clearable="true"  placeholder="请输入调动事由" />
       </el-form-item>
        <el-form-item label="原部门意见:" prop="sourceResult">
          <el-input v-model="formData.sourceResult" :clearable="true"  placeholder="请输入原部门意见" />
       </el-form-item>
        <el-form-item label="调往部门意见:" prop="toResult">
          <el-input v-model="formData.toResult" :clearable="true"  placeholder="请输入调往部门意见" />
       </el-form-item>
        <el-form-item label="调入时间:" prop="toDate">
          <el-date-picker v-model="formData.toDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="考察期:" prop="inspectionPerion">
        <el-select v-model="formData.inspectionPerion" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="附件:" prop="attachment">
          <el-input v-model="formData.attachment" :clearable="true"  placeholder="请输入附件" />
       </el-form-item>
        <el-form-item label="提交意见:" prop="submitOpinion">
          <el-input v-model="formData.submitOpinion" :clearable="true"  placeholder="请输入提交意见" />
       </el-form-item>
        <el-form-item label="OAID:" prop="oaId">
          <el-input v-model="formData.oaId" :clearable="true"  placeholder="请输入OAID" />
       </el-form-item>
        <el-form-item label="OA状态:" prop="oaStatus">
          <el-switch v-model="formData.oaStatus" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createWcStaffTransferApplication,
  updateWcStaffTransferApplication,
  findWcStaffTransferApplication
} from '@/api/employee/wcStaffTransferApplication'

defineOptions({
    name: 'WcStaffTransferApplicationForm'
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
               }],
               employmentDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               transferType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               transferStaff : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sourceDepartment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sourcePosition : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               newDepartment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               newPosition : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               transferResult : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sourceResult : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               toResult : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               toDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               inspectionPerion : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               attachment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               submitOpinion : [{
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
      const res = await findWcStaffTransferApplication({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcStaffTransferApplication
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
