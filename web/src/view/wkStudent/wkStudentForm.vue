<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="name:" prop="name">
          <el-input v-model="formData.name" :clearable="true" :placeholder="t('general.pleaseEnter')" />
       </el-form-item>
        <el-form-item label="age:" prop="age">
          <el-input v-model="formData.age" :clearable="true" :placeholder="t('general.pleaseEnter')" />
       </el-form-item>
        <el-form-item label="gender:" prop="gender">
          <el-select v-model="formData.gender" placeholder="t('general.pleaseSelect')" :clearable="true">
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">{{ t('general.save') }}</el-button>
          <el-button type="primary" @click="back">{{ t('general.back') }}</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createWkStudent,
  updateWkStudent,
  findWkStudent
} from '@/api/wkStudent'

defineOptions({
    name: 'WkStudentForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const route = useRoute()
const router = useRouter()

const type = ref('')
const genderOptions = ref([])
const formData = ref({
            name: '',
            age: '',
            gender: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWkStudent({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewkStudent
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    genderOptions.value = await getDictFunc('gender')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createWkStudent(formData.value)
               break
             case 'update':
               res = await updateWkStudent(formData.value)
               break
             default:
               res = await createWkStudent(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message:  t('general.createUpdateSuccess')
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
