<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="员工ID:" prop="staffId">
          <el-input v-model.number="formData.staffId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="合同公司:" prop="company">
          <el-input v-model="formData.company" :clearable="true"  placeholder="请输入合同公司" />
       </el-form-item>
        <el-form-item label="合同类型(1:固定期限劳动合同2:无固定期限劳动合同3:实习协议4:外包协议5:劳务派遣合同6:返聘协议7:培训协议):" prop="type">
          <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="合同起始日:" prop="startDay">
          <el-date-picker v-model="formData.startDay" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="合同到期日:" prop="endDay">
          <el-date-picker v-model="formData.endDay" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="续签次数:" prop="times">
          <el-input v-model.number="formData.times" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="合同附件:" prop="attachment">
          <el-input v-model="formData.attachment" :clearable="true"  placeholder="请输入合同附件" />
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
  createWcStaffAgreement,
  updateWcStaffAgreement,
  findWcStaffAgreement
} from '@/api/weChat/wcStaffAgreement'

defineOptions({
    name: 'WcStaffAgreementForm'
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
            staffId: 0,
            company: '',
            type: 0,
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
               }],
               company : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               startDay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               endDay : [{
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
      const res = await findWcStaffAgreement({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcStaffAgreement
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
