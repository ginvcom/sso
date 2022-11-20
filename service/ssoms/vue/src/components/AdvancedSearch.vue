<template>
  <div>
    <a-form ref="formRef" class="advanced-search" :colon="false" :model="formState" @finish="onFinish">
      <div class="advanced-search__inner" :class="{ 'is-collapsed': collapsed }">
        <div class="advanced-search__form">
          <a-row :gutter="!collapsed ? 32 : 8">
            <template v-for="(formItem, i) in formData.formItems" :key="i">
              <a-col v-show="!collapsed || formItem.advanced !== true" :span="!collapsed ? 8 : formItem.col || 8" :offset="i == 0 ? offset : 0">
                <a-form-item v-if="formItem.type == 'input'" :name="formItem.form.name" :label="formItem.form.label" :class="{ 'not-nil': formState[formItem.form.name] !== '', 'is-active': formItem.focus }">
                  <a-input v-model:value="formState[formItem.form.name]" :allow-clear="formItem.form.allowClear" @focus="onFocus(i)" @blur="onBlur(i)"></a-input>
                </a-form-item>
                <a-form-item v-else-if="formItem.type == 'select'" :name="formItem.form.name" :label="formItem.form.label" :class="{ 'not-nil': formState[formItem.form.name] !== '' && formState[formItem.form.name] !== undefined, 'is-active': formItem.focus }">
                  <a-select v-if="formItem.form" v-model:value="formState[formItem.form.name]" :options="formItem.options" :allow-clear="formItem.form.allowClear" @focus="onFocus(i)" @blur="onBlur(i)" @change="onFinish"></a-select>
                </a-form-item>
                <a-form-item v-else-if="formItem.type == 'select-extra'" :name="formItem.form.name" :label="formItem.form.label" :class="{ 'not-nil': formState[formItem.form.name] !== '' && formState[formItem.form.name] !== undefined, 'is-active': formItem.focus }">
                  <a-select v-model:value="formState[formItem.form.name]" :filter-option="filterOption" show-search :allow-clear="formItem.form.allowClear" @focus="onFocus(i)" @blur="onBlur(i)" @change="onFinish">
                    <a-select-option v-for="option in formItem.options" :key="option.value" :value="option.value" :label="option.label" :extra="option.extra">
                      <div style="display: flex; justify-content: space-between">
                        <span style="text-overflow: ellipsis; overflow: hidden">{{ option.label }}</span>
                        <span style="text-overflow: ellipsis; overflow: hidden; color: rgba(0, 0, 0, 0.45)">{{ option.extra }}</span>
                      </div>
                    </a-select-option>
                  </a-select>
                </a-form-item>
                <a-form-item v-else-if="formItem.type == 'daterange'" :name="formItem.form.name" :label="formItem.form.label" :class="{ 'not-nil': (formState[formItem.form.name] && (formState[formItem.form.name] as Array<string>).length == 2), 'is-active': formItem.focus }">
                  <a-range-picker @focus="onFocus(i)" @blur="onBlur(i)" @change="onFinish" v-model:value="formState[formItem.form.name]" valueFormat="YYYY-MM-DD" :placeholder="[]" style="width: 100%" />
                </a-form-item>
                <a-form-item v-else-if="formItem.type == 'datetimerange'" :name="formItem.form.name" :label="formItem.form.label" :class="{'not-nil': (formState[formItem.form.name] && (formState[formItem.form.name] as Array<string>).length == 2), 'is-active': formItem.focus }">
                  <a-range-picker show-time @focus="onFocus(i)" @blur="onBlur(i)" @change="onFinish" v-model:value="formState[formItem.form.name]" valueFormat="YYYY-MM-DD HH:mm:ss" :placeholder="[]" style="width: 100%" />
                </a-form-item>
              </a-col>
            </template>
            <!-- <a-col v-show="collapsed" :span="4">
              <a-button type="primary" html-type="submit">搜索</a-button>
              <a-button style="margin: 0 8px" @click="onReset">清除</a-button>
              <a v-if="formData.formItems.some(item => item.advanced === true)" style="font-size: 12px" @click="collapsed = !collapsed">
                <DownOutlined />
                高级
              </a>
            </a-col> -->
          </a-row>
        </div>
        <div class="advanced-search__actions">
          <a-row>
            <a-col :span="24">
              <a-button type="primary" html-type="submit">搜索</a-button>
              <a-button style="margin: 0 8px" @click="onReset">清除</a-button>
              <a v-if="formData.formItems.some(item => item.advanced === true)" style="font-size: 12px" @click="onCollapse">
                <UpOutlined />
                基本
              </a>
            </a-col>
          </a-row>
        </div>
      </div>
    </a-form>
  </div>
</template>
<script setup lang="ts">
import { reactive, ref } from 'vue'
import { DownOutlined, UpOutlined } from '@ant-design/icons-vue'
import type { FormInstance } from 'ant-design-vue'
import { debounce } from 'lodash-es'

export interface Option {
  label: string
  value: string | number
  extra?: string
}

export interface FormItem {
  type: string
  form: FormItemForm
  options?: Array<Option>
  syncOptions?: () => Promise<Option[]>
  col?: number
  focus?: boolean
  advanced?: boolean
}

export interface FormItemForm {
  label: string
  name: string
  allowClear: boolean
}

const collapsed = ref(true)
const formRef = ref<FormInstance>()

interface Props {
  offset?: number
  value: Record<string, string | number | Array<string> | null>
  data: Array<FormItem>
}

const props = withDefaults(defineProps<Props>(), {
  offset: 0
})

const formState = reactive<Record<string, string | number | Array<string> | null>>(props.value)

interface FormData {
  formItems: Array<FormItem>
}

const formData = reactive<FormData>({
  formItems: props.data
})

formData.formItems.forEach(async formItem => {
  if (formItem.syncOptions) {
    formItem.options = await formItem.syncOptions()
  }
})

const emit = defineEmits(['change'])

const onFinish = debounce(
  () => {
    emit('change', formState)
  },
  500,
  { leading: true, trailing: false }
)

const onReset = () => {
  formRef.value!.resetFields()
  onFinish()
}

const onCollapse = () => {
  collapsed.value = true
  let hasDifferent = false // 收起的筛选项和初始值是否改变过
  formData.formItems.forEach(formItem => {
    if (formItem.advanced === true) {
      if (formState[formItem.form.name] !== props.value[formItem.form.name]) {
        hasDifferent = true
        formState[formItem.form.name] = props.value[formItem.form.name]
      }
    }
  })
  if (hasDifferent) {
    onFinish()
  }
}

const filterOption = (input: string, opt: Option) => {
  return opt.label.toLowerCase().indexOf(input.toLowerCase()) >= 0 || opt.extra!.toLowerCase().indexOf(input.toLowerCase()) >= 0
}

const onFocus = (i: number) => {
  formData.formItems[i].focus = true
}

const onBlur = (i: number) => {
  formData.formItems[i].focus = false
}
</script>
<style scoped>
.advanced-search__inner.is-collapsed {
  display: flex;
}

.advanced-search__inner {
  display: block;
}
.advanced-search__inner.is-collapsed .advanced-search__form {
  flex: 1;
  margin-right: 8px;
}
#components-form-demo-advanced-search .ant-form {
  max-width: none;
}
#components-form-demo-advanced-search .search-result-list {
  margin-top: 16px;
  border: 1px dashed #e9e9e9;
  border-radius: 2px;
  background-color: #fafafa;
  min-height: 200px;
  text-align: center;
  padding-top: 80px;
}
[data-theme='dark'] .advanced-search {
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid #434343;
  padding: 24px;
  border-radius: 2px;
}
[data-theme='dark'] #components-form-demo-advanced-search .search-result-list {
  border: 1px dashed #434343;
  background: rgba(255, 255, 255, 0.04);
}

.advanced-search:deep() .ant-form-item {
  position: relative;
  margin-bottom: 16px;
}

.advanced-search:deep() .ant-form-item-label {
  position: absolute;
  left: 0;
  padding: 0 10px;
  pointer-events: none;
  font-size: 1em;
  transition: 0.5s;
  z-index: 22;
}

.advanced-search:deep() .ant-form-item-label > label {
  color: rgba(0, 0, 0, 0.25);
}

.advanced-search:deep() .ant-form-item.is-active .ant-form-item-label,
.advanced-search:deep() .ant-form-item.not-nil .ant-form-item-label {
  transform: scale(0.65, 0.65);
  transform-origin: 35px -100%;
  border: 1px solid #d9d9d9;
  background-color: #fff;
}

.advanced-search:deep() .ant-form-item.is-active .ant-form-item-label {
  border-color: #40a9ff;
  background-color: #e6f7ff;
}

.advanced-search:deep() .ant-form-item.is-active .ant-form-item-label > label,
.advanced-search:deep() .ant-form-item.not-nil .ant-form-item-label > label {
  letter-spacing: 0.2em;
  height: 20px;
  color: rgba(0, 0, 0, 0.85);
}

.advanced-search:deep() .ant-form-item-label > label.ant-form-item-no-colon::after {
  display: none;
}
</style>
