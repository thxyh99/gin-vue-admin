<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="成员UserID:" prop="userid">
          <el-input v-model="formData.userid" :clearable="true"  placeholder="请输入成员UserID" />
       </el-form-item>
        <el-form-item label="成员名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入成员名称" />
       </el-form-item>
        <el-form-item label="成员所属部门ID:" prop="department">
          <el-input v-model="formData.department" :clearable="true"  placeholder="请输入成员所属部门ID" />
       </el-form-item>
        <el-form-item label="排序:" prop="order">
          <el-input v-model.number="formData.order" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="职务信息ID:" prop="positionId">
          <el-input v-model.number="formData.positionId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="职务信息:" prop="position">
          <el-input v-model="formData.position" :clearable="true"  placeholder="请输入职务信息" />
       </el-form-item>
        <el-form-item label="性别(0未知1男2女):" prop="gender">
          <el-input v-model.number="formData.gender" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true"  placeholder="请输入邮箱" />
       </el-form-item>
        <el-form-item label="企业邮箱:" prop="bizMail">
          <el-input v-model="formData.bizMail" :clearable="true"  placeholder="请输入企业邮箱" />
       </el-form-item>
        <el-form-item label="表示在所在的部门内是否为部门负责人:" prop="isLeaderInDept">
          <el-input v-model="formData.isLeaderInDept" :clearable="true"  placeholder="请输入表示在所在的部门内是否为部门负责人" />
       </el-form-item>
        <el-form-item label="直属上级UserID:" prop="directLeader">
          <el-input v-model="formData.directLeader" :clearable="true"  placeholder="请输入直属上级UserID" />
       </el-form-item>
        <el-form-item label="头像:" prop="avatar">
          <el-input v-model="formData.avatar" :clearable="true"  placeholder="请输入头像" />
       </el-form-item>
        <el-form-item label="头像缩略图:" prop="thumbAvatar">
          <el-input v-model="formData.thumbAvatar" :clearable="true"  placeholder="请输入头像缩略图" />
       </el-form-item>
        <el-form-item label="座机:" prop="telephone">
          <el-input v-model="formData.telephone" :clearable="true"  placeholder="请输入座机" />
       </el-form-item>
        <el-form-item label="别名:" prop="alias">
          <el-input v-model="formData.alias" :clearable="true"  placeholder="请输入别名" />
       </el-form-item>
        <el-form-item label="地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入地址" />
       </el-form-item>
        <el-form-item label="全局唯一:" prop="openUserid">
          <el-input v-model="formData.openUserid" :clearable="true"  placeholder="请输入全局唯一" />
       </el-form-item>
        <el-form-item label="主部门:" prop="mainDepartment">
          <el-input v-model.number="formData.mainDepartment" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="扩展属性:" prop="extattr">
          <RichEdit v-model="formData.extattr"/>
       </el-form-item>
        <el-form-item label="激活状态(1已激活，2已禁用，4未激活，5退出企业):" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="员工个人二维码:" prop="qrCode">
          <el-input v-model="formData.qrCode" :clearable="true"  placeholder="请输入员工个人二维码" />
       </el-form-item>
        <el-form-item label="对外职务:" prop="externalPosition">
          <el-input v-model="formData.externalPosition" :clearable="true"  placeholder="请输入对外职务" />
       </el-form-item>
        <el-form-item label="成员对外属性:" prop="externalProfile">
          <RichEdit v-model="formData.externalProfile"/>
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
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            userid: '',
            name: '',
            department: '',
            order: 0,
            positionId: 0,
            position: '',
            gender: 0,
            email: '',
            bizMail: '',
            isLeaderInDept: '',
            directLeader: '',
            avatar: '',
            thumbAvatar: '',
            telephone: '',
            alias: '',
            address: '',
            openUserid: '',
            mainDepartment: 0,
            extattr: '',
            status: 0,
            qrCode: '',
            externalPosition: '',
            externalProfile: '',
        })
// 验证规则
const rule = reactive({
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
               department : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               order : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               positionId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               position : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               gender : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
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
