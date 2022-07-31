<template>
  <div class="password-reset">
    <a-form
      :model="formState"
      :label-col="{ span: 8 }"
      :wrapper-col="{ span: 8 }"
      autocomplete="off">
      <a-form-item
        label="个人说明"
        name="summary"
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
        <a-button type="primary">提交</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>
<script setup lang="ts">
import { reactive } from 'vue'
interface FormState {
  introduction: string
}

const formState = reactive<FormState>({
  introduction: ''
})
</script>
