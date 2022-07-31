<template>
  <div class="password-reset">
    <a-form
      :model="formState"
      :label-col="{ span: 8 }"
      :wrapper-col="{ span: 8 }"
      autocomplete="off">
      <a-form-item
        label="原密码"
        name="username"
        :rules="[{ required: true, message: '请输入现在的密码!' }]"
      >
        <a-input-password v-model:value="formState.oldPassword" placeholder="请输入现在的密码" />
      </a-form-item>

      <a-form-item
        label="新密码"
        name="password"
        :rules="[{ required: true, message: '请输入新密码!' }]"
      >
        <a-input-password v-model:value="formState.password" placeholder="请输入新密码" />
      </a-form-item>

      <a-form-item
        label="新密码确认"
        name="confirmPassword"
        :rules="[{ required: true, message: '请再次输入新密码!' }]"
      >
        <a-input-password v-model:value="formState.confirmPassword" placeholder="两次输入密码保持一致" />
      </a-form-item>

      <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
        <a-button type="primary">提交</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>
<script setup lang="ts">
import { reactive } from 'vue'
interface FormState {
  oldPassword: string
  password: string
  confirmPassword: string
}

const formState = reactive<FormState>({
  oldPassword: '',
  password: '',
  confirmPassword: '',
})
</script>
