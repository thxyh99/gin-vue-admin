<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户ID(SSO):" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="企微成员UserID:" prop="userid">
          <el-input v-model="formData.userid" :clearable="true"  placeholder="请输入企微成员UserID" />
       </el-form-item>
        <el-form-item label="紧急联系人姓名:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入紧急联系人姓名" />
       </el-form-item>
        <el-form-item label="联系人关系(1:父母 2:配偶 3:子女 0:其他):" prop="relationship">
          <el-input v-model.number="formData.relationship" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="联系人电话:" prop="mobile">
          <el-input v-model="formData.mobile" :clearable="true"  placeholder="请输入联系人电话" />
       </el-form-item>
        <el-form-item label="联系人常住地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入联系人常住地址" />
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
  createWcStaffContact,
  updateWcStaffContact,
  findWcStaffContact
} from '@/api/weChat/wcStaffContact'

defineOptions({
    name: 'WcStaffContactForm'
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
            userId: 0,
            userid: '',
            name: '',
            relationship: 0,
            mobile: '',
            address: '',
        })
// 验证规则
const rule = reactive({
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userid : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               relationship : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mobile : [{
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
      const res = await findWcStaffContact({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcStaffContact
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
               res = await createWcStaffContact(formData.value)
               break
             case 'update':
               res = await updateWcStaffContact(formData.value)
               break
             default:
               res = await createWcStaffContact(formData.value)
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
