<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="转正标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入转正标题" />
       </el-form-item>
        <el-form-item label="员工ID:" prop="staffId">
          <el-input v-model.number="formData.staffId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="入职日期:" prop="employmentDate">
          <el-date-picker v-model="formData.employmentDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="所属部门:" prop="jobDepartment">
        <el-select v-model="formData.jobDepartment" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="原职位:" prop="sourcePosition">
        <el-select v-model="formData.sourcePosition" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="原职级:" prop="sourceLevel">
        <el-select v-model="formData.sourceLevel" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他):" prop="tryPeriod">
        <el-select v-model="formData.tryPeriod" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="是否是部门负责人:" prop="isMananger">
          <el-switch v-model="formData.isMananger" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="附件:" prop="attachment">
          <el-input v-model="formData.attachment" :clearable="true"  placeholder="请输入附件" />
       </el-form-item>
        <el-form-item label="个人自我鉴定:" prop="selfOpinion">
          <el-input v-model="formData.selfOpinion" :clearable="true"  placeholder="请输入个人自我鉴定" />
       </el-form-item>
        <el-form-item label="用人部门意见:" prop="tryOpinion">
          <el-input v-model="formData.tryOpinion" :clearable="true"  placeholder="请输入用人部门意见" />
       </el-form-item>
        <el-form-item label="转正后工作职责:" prop="newResponsibilities">
        <el-select v-model="formData.newResponsibilities" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [255]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="转正时间:" prop="passDate">
        <el-select v-model="formData.passDate" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="转正后职位:" prop="newPosition">
        <el-select v-model="formData.newPosition" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="转正后职级:" prop="newJoblevel">
        <el-select v-model="formData.newJoblevel" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
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
  createWcStaffPassApplication,
  updateWcStaffPassApplication,
  findWcStaffPassApplication
} from '@/api/employee/wcStaffPassApplication'

defineOptions({
    name: 'WcStaffPassApplicationForm'
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
               }],
               employmentDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               jobDepartment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sourcePosition : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sourceLevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               tryPeriod : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               isMananger : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               attachment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               selfOpinion : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               tryOpinion : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               newResponsibilities : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               passDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               newPosition : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               newJoblevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               submitOpinion : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               oaId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               oaStatus : [{
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
      const res = await findWcStaffPassApplication({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcStaffPassApplication
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
