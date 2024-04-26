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
        <el-form-item label="成员名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入成员名称" />
       </el-form-item>
        <el-form-item label="别名:" prop="alias">
          <el-input v-model="formData.alias" :clearable="true"  placeholder="请输入别名" />
       </el-form-item>
        <el-form-item label="职务信息ID:" prop="positionId">
          <el-input v-model.number="formData.positionId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别(0未知1男2女):" prop="gender">
          <el-input v-model.number="formData.gender" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否领导(1:是 0:否):" prop="isLeader">
          <el-input v-model.number="formData.isLeader" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手机:" prop="mobile">
          <el-input v-model="formData.mobile" :clearable="true"  placeholder="请输入手机" />
       </el-form-item>
        <el-form-item label="座机:" prop="telephone">
          <el-input v-model="formData.telephone" :clearable="true"  placeholder="请输入座机" />
       </el-form-item>
        <el-form-item label="个人邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true"  placeholder="请输入个人邮箱" />
       </el-form-item>
        <el-form-item label="地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入地址" />
       </el-form-item>
        <el-form-item label="企业邮箱:" prop="bizMail">
          <el-input v-model="formData.bizMail" :clearable="true"  placeholder="请输入企业邮箱" />
       </el-form-item>
        <el-form-item label="视频号:" prop="videoId">
          <el-input v-model="formData.videoId" :clearable="true"  placeholder="请输入视频号" />
       </el-form-item>
        <el-form-item label="英文名:" prop="nameEn">
          <el-input v-model="formData.nameEn" :clearable="true"  placeholder="请输入英文名" />
       </el-form-item>
        <el-form-item label="激活状态(1已激活，2已禁用，4未激活，5退出企业):" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否关注微信插件:" prop="isAttention">
          <el-input v-model.number="formData.isAttention" :clearable="true" placeholder="请输入" />
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
  createWcStaff,
  updateWcStaff,
  findWcStaff
} from '@/api/weChat/wcStaff'

defineOptions({
    name: 'WcStaffForm'
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
            alias: '',
            positionId: 0,
            gender: 0,
            isLeader: 0,
            mobile: '',
            telephone: '',
            email: '',
            address: '',
            bizMail: '',
            videoId: '',
            nameEn: '',
            status: 0,
            isAttention: 0,
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
               gender : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               isLeader : [{
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
      const res = await findWcStaff({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcStaff
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
