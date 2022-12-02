<template>
  <div class="password-reset">
    <a-form
      :model="formState"
      :label-col="{ span: 8 }"
      :wrapper-col="{ span: 8 }"
      autocomplete="off"
      @finish="onFinish">
      <a-form-item
        label="个人说明"
        name="introduction"
        :rules="[{ required: true, message: '请用一段话介绍自己!' }]"
      >
        <a-textarea
          v-model:value="formState.introduction"
          placeholder="填写职业技能、擅长的事情、喜欢的事情等"
          :rows="5"
          show-count
          :maxlength="100"
        />
      </a-form-item>
      <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
        <a-button type="primary" html-type="submit">提交</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>
<script setup lang="ts">
import { message } from 'ant-design-vue'
import { reactive } from 'vue'
import {
  infoEdit
} from '../api/ssoms'
export interface UserForm {
	uuid?: string
	name: string
	mobile: string
	password?: string
	avatar: string
	gender: number
	birth: string
	introduction: string
	status: number
}

interface FormState {
  introduction: string
}

const props = withDefaults(defineProps<UserForm>(), {
  introduction: ''
})

const formState = reactive<FormState>({
  introduction: props.introduction
})

const emit = defineEmits(['change'])

const onFinish = (values: any) => {
  infoEdit(values).then(() => {
    message.success('用户信息修改成功')
    emit('change', values)
  })
}
</script>
