<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="调动标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入调动标题" />
       </el-form-item>
        <el-form-item label="员工ID:" prop="staffId">
          <el-input v-model.number="formData.staffId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="生效日期:" prop="effectiveDate">
          <el-date-picker v-model="formData.effectiveDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="调动前部门:" prop="sourceDepartment">
        <el-select v-model="formData.sourceDepartment" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="调动后部门:" prop="newDepartment2">
        <el-select v-model="formData.newDepartment2" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="调动前岗位:" prop="sourcePosition">
        <el-select v-model="formData.sourcePosition" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="调动前岗位:" prop="newPosition">
        <el-select v-model="formData.newPosition" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="调动前职级:" prop="sourceJoblevel">
        <el-select v-model="formData.sourceJoblevel" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="调动后职级:" prop="newJoblevel">
        <el-select v-model="formData.newJoblevel" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="调动前上级:" prop="sourceManager">
        <el-select v-model="formData.sourceManager" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="调动后上级:" prop="newManager">
        <el-select v-model="formData.newManager" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="备注:" prop="memo">
          <el-input v-model="formData.memo" :clearable="true"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item label="附件:" prop="attachment">
          <el-input v-model="formData.attachment" :clearable="true"  placeholder="请输入附件" />
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
  createWcStaffAjustlevelApplication,
  updateWcStaffAjustlevelApplication,
  findWcStaffAjustlevelApplication
} from '@/api/employee/wcStaffAjustlevelApplication'

defineOptions({
    name: 'WcStaffAjustlevelApplicationForm'
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
            effectiveDate: new Date(),
            memo: '',
            attachment: '',
            oaId: '',
            oaStatus: false,
        })
// 验证规则
const rule = reactive({
               effectiveDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sourceDepartment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               newDepartment2 : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sourcePosition : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               newPosition : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sourceJoblevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               newJoblevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sourceManager : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               newManager : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               memo : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               attachment : [{
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
      const res = await findWcStaffAjustlevelApplication({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcStaffAjustlevelApplication
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
               res = await createWcStaffAjustlevelApplication(formData.value)
               break
             case 'update':
               res = await updateWcStaffAjustlevelApplication(formData.value)
               break
             default:
               res = await createWcStaffAjustlevelApplication(formData.value)
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
