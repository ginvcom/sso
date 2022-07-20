<template>
  <div class="content-header">
    <div>
      <h1>菜单 & 操作</h1>
    </div>
    <div class="content-header__actions">
      <a-button type="primary" @click="initAdd">添加对象</a-button>
    </div>
  </div>
  <div class="system-params">
  <a-select
      v-model:value="state.params.topKey"
      style="width: 100%"
      size="large"
      placeholder="select one country"
      option-label-prop="children"
    >
      <a-select-option value="china" label="China">
        <img role="img" aria-label="China" src="https://dss2.bdstatic.com/8_V1bjqh_Q23odCf/pacific/1993343984.jpg" />
        &nbsp;&nbsp;China (中国)
      </a-select-option>
      <a-select-option value="usa" label="USA">
        <img role="img" aria-label="China" src="https://dss2.bdstatic.com/8_V1bjqh_Q23odCf/pacific/1993343984.jpg" />
        &nbsp;&nbsp;USA (美国)
      </a-select-option>
      <a-select-option value="japan" label="Japan">
        <img role="img" aria-label="China" src="https://dss2.bdstatic.com/8_V1bjqh_Q23odCf/pacific/1993343984.jpg" />
        &nbsp;&nbsp;Japan (日本)
      </a-select-option>
      <a-select-option value="korea" label="Korea">
        <img role="img" aria-label="China" src="https://dss2.bdstatic.com/8_V1bjqh_Q23odCf/pacific/1993343984.jpg" />
        &nbsp;&nbsp;Korea (韩国)
      </a-select-option>
    </a-select>
  </div>
  <a-table
  :dataSource="respState.list"
  :columns="columns"
  :pagination="false">
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'objectName'">
        <span style="margin-left: 10px">{{record.objectName}}</span>
      </template>
      <template v-if="column.key === 'gender'">
        <span v-if="record.gender == 1">男</span>
        <span v-else-if="record.gender == 2">女</span>
        <span v-else-if="record.gender == 3">未知</span>
      </template>
      <template v-if="column.key === 'actions'">
        <a @click="initEdit(record.uuid)">编辑</a>
        <a-divider type="vertical" />
        <a-popconfirm
          title="确定要删除该对象吗?"
          ok-text="Yes"
          cancel-text="No"
          @confirm="onDelete(record.uuid)">
          <a>删除</a>
        </a-popconfirm>
      </template>
    </template>
  </a-table>
  <a-modal
    width="720px"
    v-model:visible="formState.visible"
    :title="formState.type == 'add' ? '新增对象' : '修改对象'"
    :loading="formState.loading"
    @cancel="onCancel"
    @ok="onSubmit">
    <a-form layout="vertical" ref="modalFormRef" :model="formState.form">
      <a-row :gutter="24">
        <a-col :span="12">
          <a-form-item label="菜单名" name="form.objectName">
            <a-input v-model:value="formState.form.objectName" placeholder="姓名" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item label="菜单path" name="form.key">
            <a-input v-model:value="formState.form.key" placeholder="菜单path" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item label="菜单类型" name="form.type">
            <a-radio-group v-model:value="formState.form.type">
              <a-radio :value="1">菜单</a-radio>
              <a-radio :value="2">菜单组</a-radio>
            </a-radio-group>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item label="域名" name="form.domain">
            <a-input v-model:value="formState.form.domain" placeholder="13位手机号" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item label="状态" name="form.status">
            <a-switch v-model:checked="formState.form.status" :checked-value="1" :unChecked-value="0" checked-children="启用" un-checked-children="停用" />
          </a-form-item>
        </a-col>
      </a-row>
    </a-form>
  </a-modal>
</template>
<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import {
  objectList,
  objectDetail,
  addObject,
  updateObject,
  deleteObject,
  ObjectListReply,
  ObjectForm
} from '../api/ssoms'
import type { FormInstance, UploadChangeParam, UploadProps } from 'ant-design-vue'
import { message } from 'ant-design-vue'

/**
 * 这是表格的列定义
 * 没有动态变化，所以不使用reactive
 */
const columns = [
  {
    title: 'uuid',
    dataIndex: 'uuid',
    key: 'uuid',
    width: '200px',
  },
  {
    title: '对象名称',
    dataIndex: 'objectName',
    key: 'objectName'
  },
  {
    title: '类型',
    dataIndex: 'type',
    key: 'type'
  },
  {
    title: 'key',
    dataIndex: 'key',
    key: 'key'
  },
  {
    title: '排序',
    dataIndex: 'sort',
    key: 'sort'
  },
  {
    title: '操作',
    key: 'actions',
    align: 'center',
    width: '120px'
  }
]


/**
 * 这是列表的入参
 * 没有动态变化，所以不使用reactive
 */
const state = reactive({
  loading: false, // 列表是否加载完成
  params: {
    topKey: 'china',
    objectName: undefined,
    key: undefined
  }
})

/**
 * 这是列表响应
 */
const respState = reactive<ObjectListReply>({
  list: []
})

onMounted(() => {
  getList()
})

/**
 * 获取列表
 */
const getList = () => {
  state.loading = true
  objectList(state.params).then((data: ObjectListReply) => {
    respState.list = data.list
  }).finally(() => {
    state.loading = false
  })
}

/**
 * 删除对象
 */
const onDelete = (uuid: string) => {
  deleteObject({ uuid }).then(() => {
    message.success('删除对象操作成功')
    getList()
  })
}

const modalFormRef = ref<FormInstance>()

interface Form {
  type: string
  visible: boolean,
  loading: boolean,
  form: ObjectForm
}

const formState = reactive<Form>({
  visible: false,
  loading: false,
  type: 'add',
  form: {
    uuid: '',
    objectName: '',
    domain: '',
    key: '',
    icon: '',
    type: 1,
    pUUID: '',
    status: 1,
    sort: 0,
    topKey: ''
  }
})

const initAdd = () => {
  // resetFileds初始化无效，手动初始化
  formState.form = {
    uuid: '',
    objectName: '',
    domain: '',
    key: '',
    icon: '',
    type: 1,
    pUUID: '',
    status: 1,
    sort: 0,
    topKey: ''
  }
  formState.type = 'add'
  formState.visible = true
}


/**
 * 获取对象详情, 用于编辑
 * @param uuid
 */
const initEdit = (uuid: string) => {
  objectDetail({ uuid }).then((data: ObjectForm) => {
    formState.form = data
    formState.type = 'edit'
    formState.visible = true
  })
}

/**
 * 新增或修改对象
 */
const onSubmit = () =>{
  modalFormRef.value?.validate().then(() => {
    formState.loading = true
    if (formState.type === 'add') {
      addObject(formState.form).then(() => {
        message.success('新增对象成功')
        formState.visible = false
        modalFormRef.value?.resetFields()
        getList()
      }).finally(() => {
        formState.loading = false
      })
    } else {
      updateObject(formState.form).then(() => {
        message.success('修改对象成功')
        formState.visible = false
        modalFormRef.value?.resetFields()
        getList()
      }).finally(() => {
        formState.loading = false
      })
    }
  }).catch((err) => {
    console.log(err)
  })
}

/**
 * 关闭modelForm，重置表单
 */
const onCancel = () => {
  modalFormRef.value?.resetFields()
}
</script>

<style scoped>
.avatar-uploader:deep() > .ant-upload {
  width: 120px;
  height: 120px;
}
.system-params{
  margin: 20px 0;
}
</style>
