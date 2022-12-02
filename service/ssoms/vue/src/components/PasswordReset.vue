<template>
  <div class="password-reset">
    <a-form
      :model="formState.form"
      :label-col="{ span: 8 }"
      :wrapper-col="{ span: 8 }"
      autocomplete="off"
      @finish="onFinish"
      >
      <a-form-item
        label="原密码"
        name="oldPassword"
        :rules="[{ required: true, message: '请输入现在的密码!' }]"
      >
        <a-input-password v-model:value="formState.form.oldPassword" placeholder="请输入现在的密码" />
      </a-form-item>

      <a-form-item
        label="新密码"
        name="password"
        :rules="[{ required: true, message: '请输入新密码!' }]"
      >
        <a-input-password v-model:value="formState.form.password" placeholder="请输入新密码" />
      </a-form-item>

      <a-form-item
        label="新密码确认"
        name="confirmPassword"
        :rules="[{ required: true, message: '请再次输入新密码!' }]"
      >
        <a-input-password v-model:value="formState.form.confirmPassword" placeholder="两次输入密码保持一致" />
      </a-form-item>

      <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
        <a-button type="primary" html-type="submit">提交</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>
<script setup lang="ts">
import { reactive } from 'vue'
import { message } from 'ant-design-vue'
import {
  PasswordResetReq,
  passwordReset
} from '../api/ssoms'
interface FormState {
  form: PasswordResetReq
  loading: boolean
}

const formState = reactive<FormState>({
  loading: false,
  form: {
    oldPassword: '',
    password: '',
    confirmPassword: ''
  }
})

const onFinish = (values: PasswordResetReq) => {
    formState.loading = true
    passwordReset(values).then(() => {
      message.success('密码修改成功')
      formState.form = {
        oldPassword: '',
        password: '',
        confirmPassword: ''
      }
    }).finally(() => {
      formState.loading = false
    })
}
</script>
