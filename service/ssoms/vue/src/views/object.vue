<template>
  <layout>
    <div class="content-header">
      <div>
        <h1>菜单 & 操作</h1>
      </div>
      <div class="content-header__actions">
        <a-button type="primary" @click="initAddMenu()">
          <template #icon><plus-outlined /></template>
          添加菜单
        </a-button>
        <a-button @click="onExport">
          <template #icon><download-outlined /></template>
          导出菜单
        </a-button>
        <a-button @click="onImport">
          <template #icon><upload-outlined /></template>
          导入菜单
        </a-button>
      </div>
    </div>
    <div class="system-params">
      <a-form :model="state.params" name="params">
        <div class="object__form">
          <div class="object__current-system">
            <div class="object__current-value">
              <a-image
                :width="48"
                :height="48"
                :src="ossConfig.ginvdoc.domain + state.currentSystem.icon"
                :preview="false"
                fallback="data:image/png;base64,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"
              />
              <p class="object__current-system-name">{{state.currentSystem.name}}</p>
            </div>
            <a-button @click="initChangeSystem">
              <template #icon><appstore-outlined /></template>
              切换系统
            </a-button>
          </div>
          <a-divider type="vertical" style="height: 48px;margin: 0 48px;" />
          <div class="object__form-right">
            <advanced-search :value="searchForm" :data="searchFormData" @change="onSearch" />
          </div>
        </div>
      </a-form>
    </div>
    <a-table
    :loading="state.loading"
    rowKey="uuid"
    :dataSource="respState.list"
    :columns="columns"
    v-model:expandedRowKeys="state.expandedRowKeys"
    :pagination="false">
      <template #bodyCell="{ column, record }">
        <template v-if="column.dataIndex === 'objectName'">
          <template
            v-for="(fragment, i) in record.objectName
              .toString()
              .split(new RegExp(`(?<=${searchForm.objectName})|(?=${searchForm.objectName})`, 'i'))"
          >
            <b
              v-if="fragment.toLowerCase() === searchForm.objectName.toLowerCase()"
              :key="i"
              style="color: #f50"
            >
              {{ fragment }}
            </b>
            <template v-else>{{ fragment }}</template>
          </template>
        </template>
        <template v-if="column.dataIndex === 'type'">
          <template v-if="record.type == 2">
          <span v-if="record.subType == 1">菜单</span>
          <span v-else-if="record.subType == 2">菜单组</span>
          <span v-else-if="record.subType == 3">隐藏菜单</span>
          </template>
          <template v-else-if="record.type == 3">
            <span>操作</span>
          </template>
        </template>
        <template v-if="column.dataIndex === 'key'">
          <template v-if="record.type == 3">
            <span v-if="record.subType == 1" class="method-get m-r-15">GET</span>
            <span v-else-if="record.subType == 2" class="method-post m-r-5">POST</span>
            <span v-else-if="record.subType == 3" class="method-put m-r-15">PUT</span>
            <span v-else-if="record.subType == 4" class="method-patch m-r-15">PAT</span>
            <span v-else-if="record.subType == 5" class="method-delete m-r-15">DEL</span>
          </template>
          <template
            v-for="(fragment, i) in record.key
              .toString()
              .split(new RegExp(`(?<=${searchForm.key})|(?=${searchForm.key})`, 'i'))"
          >
            <b
              v-if="fragment.toLowerCase() === searchForm.key.toLowerCase()"
              :key="i"
              style="color: #f50"
            >
              {{ fragment }}
            </b>
            <template v-else>{{ fragment }}</template>
          </template>
        </template>
        <template v-if="column.dataIndex === 'status'">
          <span v-if="record.status == 1"><a-badge status="success" /> 启用</span>
          <span v-else-if="record.status == 0"><a-badge status="error" />停用</span>
        </template>
        <template v-if="column.dataIndex === 'actions'">
          <a v-if="record.type == 3" disabled>+子菜单</a>
          <a v-else @click="initAddMenu(record.uuid)">+子菜单</a>
          <a-divider type="vertical" />
          <a v-if="record.type == 3" disabled>+操作</a>
          <a v-else  @click="initAddApi(record.uuid)">+操作</a>
          <a-divider type="vertical" />
          <a v-if="record.type == 2" @click="initEditMenu(record.uuid)">编辑</a>
          <a v-else @click="initEditApi(record.uuid)">编辑</a>
          <a-divider type="vertical" />
          <a-popconfirm
            title="确定要删除该对象吗?"
            placement="topRight"
            @confirm="onDelete(record.uuid)">
            <a>删除</a>
          </a-popconfirm>
        </template>
      </template>
    </a-table>
    <a-modal
      width="720px"
      v-model:visible="systemFormState.visible"
      title="切换系统"
      :footer="false"
      @cancel="onSystemFormCancel">
      <div class="object__select">
        <a-row :gutter="24">
          <a-col v-for="obj in state.systems" :key="obj.uuid" :span="8">
            <div
            class="object__current-value object__option"
            :class="{ 'is-active': state.params.topKey == obj.key }"
            @click="onChangeSystem(obj)"
            >
              <a-image
                :width="48"
                :height="48"
                :src="ossConfig.ginvdoc.domain + obj.icon"
                :preview="false"
                fallback="data:image/png;base64,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"
              />
              <p class="object__current-system-name">{{obj.objectName}}</p>
            </div>
          </a-col>
        </a-row>
      </div>
    </a-modal>
    <a-modal
      width="620px"
      v-model:visible="formState.visible"
      :loading="formState.loading"
      :maskClosable="false"
      :title="formState.type == 'add' ? '添加菜单' : '修改菜单'"
      @cancel="onCancel"
      @ok="onSubmit">
      <a-form layout="vertical" ref="modalFormRef" :model="formState.form">
        <a-row :gutter="24">
          <a-col :span="12">
            <a-form-item label="菜单名" name="objectName" :rules="[{ required: true, message: '请输入菜单名' }]">
              <a-input v-model:value="formState.form.objectName" placeholder="菜单名" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="菜单路径" name="key" :rules="[{ required: true, message: '请输入菜单路径' }]">
              <a-input v-model:value="formState.form.key" placeholder="菜单路径" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="父级菜单" name="pUUID">
              <a-tree-select
                v-model:value="formState.form.pUUID"
                show-search
                style="width: 100%"
                :height="400"
                placeholder="请选择父级菜单"
                allow-clear
                :tree-line="{ showLeafIcon: false }"
                tree-default-expand-all
                :tree-data="state.menus"
              >
              </a-tree-select>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="菜单类型" name="subType" :rules="[{ required: true, message: '请选择菜单类型' }]">
              <a-radio-group v-model:value="formState.form.subType">
                <a-radio :value="1">菜单</a-radio>
                <a-radio :value="2">菜单组</a-radio>
                <a-radio :value="3">隐藏菜单</a-radio>
              </a-radio-group>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="标识符" name="identifier">
              <a-input v-model:value="formState.form.identifier" placeholder="标识" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="icon图标" name="icon">
              <a-input v-model:value="formState.form.icon" placeholder="icon图标" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="排序值" name="sort">
              <a-input-number v-model:value="formState.form.sort" placeholder="排序值小的靠前" style="width: 100%" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="状态" name="status" :rules="[{ required: true, message: '请选择菜单状态' }]">
              <a-switch v-model:checked="formState.form.status" :checked-value="1" :unChecked-value="0" checked-children="启用" un-checked-children="停用" />
            </a-form-item>
          </a-col>
        </a-row>
      </a-form>
    </a-modal>
    <a-modal
      width="720px"
      v-model:visible="formState.apiVisible"
      :loading="formState.loading"
      :maskClosable="false"
      :title="formState.type == 'add' ? '添加操作' : '修改操作'"
      @cancel="onCancel"
      @ok="onSubmit">
      <a-form layout="vertical" ref="modalFormRef" :model="formState.form">
        <a-row :gutter="24">
          <a-col :span="24">
            <a-form-item label="操作请求" name="key" :rules="[{ required: true, message: '请输入操作请求' }]">
              <a-input-group compact>
              <a-select
                  v-model:value="formState.form.subType"
                  placeholder="select one country"
                  option-label-prop="children"
                  style="width: 20%"
                >
                  <a-select-option :value="1" label="China">
                    <span class="method-get">GET</span>
                  </a-select-option>
                  <a-select-option :value="2" label="USA">
                    <span class="method-post">POST</span>
                  </a-select-option>
                  <a-select-option :value="3" label="Japan">
                    <span class="method-put">PUT</span>
                  </a-select-option>
                  <a-select-option :value="4" label="Japan">
                    <span class="method-patch">PATCH</span>
                  </a-select-option>
                  <a-select-option :value="5" label="Korea">
                    <span class="method-delete">DELETE</span>
                  </a-select-option>
                </a-select>
                <a-input v-model:value="formState.form.key" style="width: 80%" placeholder="接口路径，“/”起始" />
              </a-input-group>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="操作名" name="objectName" :rules="[{ required: true, message: '请输入操作名' }]">
              <a-input v-model:value="formState.form.objectName" placeholder="操作名" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="所属菜单" name="pUUID" :rules="[{ required: true, message: '请选择所属菜单' }]">
              <a-tree-select
                v-model:value="formState.form.pUUID"
                show-search
                style="width: 100%"
                :height="400"
                placeholder="请选择所属菜单"
                allow-clear
                :tree-line="{ showLeafIcon: false }"
                tree-default-expand-all
                :tree-data="state.apiParentMenus"
              >
              </a-tree-select>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="状态" name="status" :rules="[{ required: true, message: '请选择菜单状态' }]">
              <a-switch v-model:checked="formState.form.status" :checked-value="1" :unChecked-value="0" checked-children="启用" un-checked-children="停用" />
            </a-form-item>
          </a-col>
        </a-row>
      </a-form>
    </a-modal>
    <a-modal
      v-model:visible="importState.visible"
      :closable="importState.percent == 100"
      :maskClosable="false"
      :bodyStyle="{padding:0}"
      title="导入菜单&操作"
      :footer="null"
      @cancel="onImportCancel"
      >
      <div class="import-log">
        <a-progress style="padding:0 8px 0 20px" :percent="importState.percent" stroke-color="#52c41a" />
        <div class="import-log__content" ref="importLogContentRef">
          <div v-for="(item, i) in importState.data" :key="i">
            导入 {{item.uuid}} {{item.objectName}}
            <span v-if="item.status == 'ready'" class="import-ready">准备中<ellipsis-outlined /></span>
            <span v-else-if="item.status == 'doing'" class="import-doing">进行中<line-outlined spin /></span>
            <span v-else-if="item.status == 'success'" class="import-success">已成功<exclamation-outlined /></span>
            <span v-else-if="item.status == 'ignore'" class="import-ignore">已忽略<exclamation-outlined /></span>
            <span v-else-if="item.status == 'failed'" class="import-failed">已失败<exclamation-outlined /></span>
            <span>{{item.msg}}</span>
          </div>
        </div>
        <div class="log__content_statistics">
          <p>
            结果: 共{{importState.total}}条, 已处理{{importState.data.filter(item => item.status != 'ready' && item.status != 'doing' ).length}}条,
            成功<span class="import-success">{{importState.statistics.success}}</span>条,
            忽略<span class="import-ignore">{{importState.statistics.ignore}}</span>条,
            失败<span class="import-failed">{{importState.statistics.failed}}</span>条,
            耗时{{importState.timeCount.toFixed(2)}}秒。
          </p>
        </div>
      </div>
    </a-modal>
  </layout>
</template>
<script setup lang="ts">
import { onBeforeMount, reactive, ref, createVNode, nextTick, onBeforeUnmount } from 'vue'
import AdvancedSearch from '@/components/AdvancedSearch.vue'
import { DownloadOutlined, ExclamationCircleOutlined, EllipsisOutlined, ExclamationOutlined, LineOutlined, CloseOutlined, UploadOutlined, PlusOutlined, AppstoreOutlined } from '@ant-design/icons-vue'
// Object 是js的关键字, 别名处理一下
import {
  objectList,
  objectDetail,
  addObject,
  importObject,
  menuOptions,
  MenuOptionsReply,
  MenuOption,
  updateObject,
  deleteObject,
  ObjectListReqParams,
  ObjectListReply,
  ObjectForm,
  Object as Obj
} from '../api/ssoms'
import { ossConfig } from '@/config'
import { FormInstance, message, Modal } from 'ant-design-vue'
import * as XLSX from 'xlsx'

const searchFormData = [
  {
    type: 'input',
    col: 12,
    form: {
      label: '菜单路径',
      name: 'key',
      allowClear: true
    }
  },
  {
    type: 'input',
    col: 12,
    form: {
      label: '菜单名称',
      name: 'objectName',
      allowClear: true
    }
  }
]

const searchForm = reactive<FilterParams>({ key: '', objectName: '' })

/**
 * 这是表格的列定义
 * 没有动态变化，所以不使用reactive
 */
const columns = [
  {
    title: 'uuid',
    dataIndex: 'uuid',
    width: '200px',
  },
  {
    title: '菜单名称 / 操作名称',
    dataIndex: 'objectName'
  },
  {
    title: '类型',
    dataIndex: 'type'
  },
  {
    title: '菜单路径 / 操作uri',
    dataIndex: 'key'
  },
  {
    title: '状态',
    dataIndex: 'status'
  },
  {
    title: '排序',
    dataIndex: 'sort'
  },
  {
    title: '操作',
    dataIndex: 'actions',
    align: 'center',
    width: '230px'
  }
]

interface System {
  name: string
  icon: string
}

interface FilterParams {
  objectName: string
  key: string
}

interface State {
  loading: boolean
  currentSystem: System

  expandedRowKeys: Array<string>
  systems: Array<Obj>
  menus: Array<MenuOption>
  apiParentMenus: Array<MenuOption>
  params: ObjectListReqParams
  filterParams: FilterParams
}


/**
 * 这是列表的入参
 * 没有动态变化，所以不使用reactive
 */
const state = reactive<State>({
  loading: false, // 列表是否加载完成
  currentSystem: {
    icon: '',
    name: ''
  },
  expandedRowKeys: [],
  systems: [],
  menus: [],
  apiParentMenus: [],
  filterParams: {
    objectName: '',
    key: ''
  },
  params: {
    topKey: '',
    objectName: '',
    key: ''
  }
})

/**
 * 这是列表响应
 */
const respState = reactive<ObjectListReply>({
  list: []
})

onBeforeMount(() => {
  getSystemOptions()
})

const getSystemOptions = () => {
  objectList({ topKey: '', key: '', objectName: '' }).then((data: ObjectListReply) => {
    state.systems = data.list
    if (state.systems.length > 0) {
      const firstSystem = state.systems[0]
      state.currentSystem = { name: firstSystem.objectName, icon: firstSystem.icon }
      state.params.topKey = firstSystem.key
      getList()
    }
  }).finally(() => {
    state.loading = false
  })
}

let srcList:Array<Obj> = []

/**
 * 获取列表
 */
const getList = () => {
  state.loading = true
  objectList(state.params).then((data: ObjectListReply) => {
    srcList = data.list
    respState.list = data.list
    state.expandedRowKeys = getAllChildrenKeys(srcList, [])
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
  apiVisible: boolean,
  apiLoading: boolean,
  form: ObjectForm
}

const formState = reactive<Form>({
  visible: false,
  loading: false,
  apiVisible: false,
  apiLoading: false,
  type: 'add',
  form: {
    uuid: '',
    objectName: '',
    identifier: '',
    key: '',
    icon: '',
    type: 2,
    subType: 1,
    pUUID: '',
    status: 1,
    sort: 0,
    topKey: ''
  }
})

const initAddMenu = (pUUID?: string) => {
  menuOptions({ excludeHide: true }).then(data => {
    state.menus = data.list
    formState.form = {
      uuid: '',
      objectName: '',
      identifier: '',
      key: '',
      icon: '',
      type: 2,
      subType: 1,
      pUUID: pUUID || '',
      status: 1,
      sort: 0,
      topKey: state.params.topKey
    }
    formState.type = 'add'
    formState.visible = true
  })
}

/**
 * 获取对象详情, 用于编辑
 * @param uuid
 */
const initEditMenu = (uuid: string) => {
  objectDetail({ uuid }).then(data => {
    menuOptions({ excludeHide: true }).then(menudata => {
      state.menus = menudata.list
      formState.form = data
      formState.type = 'edit'
      formState.visible = true
    })
  })
}

const initEditApi = (uuid: string) => {
  objectDetail({ uuid }).then(data => {
    menuOptions({ excludeHide: false }).then(menudata => {
      state.apiParentMenus = menudata.list
      formState.form = data
      formState.type = 'edit'
      formState.apiVisible = true
    })
  })
}

const initAddApi = (pUUID?: string) => {
  menuOptions({ excludeHide: false }).then(data => {
    state.apiParentMenus = data.list
    formState.form = {
      uuid: '',
      objectName: '',
      identifier: '',
      key: '',
      icon: '',
      type: 3,
      subType: 1,
      pUUID: pUUID || '',
      status: 1,
      sort: 0,
      topKey: state.params.topKey
    }
    formState.type = 'add'
    formState.apiVisible = true
  })
}

/**
 * 新增或修改菜单
 */
const onSubmit = () =>{
  modalFormRef.value?.validate().then(() => {
    formState.loading = true
    if (formState.type === 'add') {
      addObject(formState.form).then(() => {
        message.success('新增菜单成功')
        formState.visible = false
        formState.apiVisible = false
        modalFormRef.value?.resetFields()
        getList()
      }).finally(() => {
        formState.apiLoading = false
        formState.loading = false
      })
    } else {
      updateObject({ uuid: formState.form.uuid! }, formState.form).then(() => {
        message.success('修改菜单成功')
        formState.visible = false
        formState.apiVisible = false
        modalFormRef.value?.resetFields()
        getList()
      }).finally(() => {
        formState.apiLoading = false
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

const systemFormState = reactive({
  visible: false //
})

const initChangeSystem = () => {
  systemFormState.visible = true
}

const onSystemFormCancel = () => {
  systemFormState.visible = false
}

const onChangeSystem = (obj: Obj) => {
  if (state.params.topKey === obj.key) {
    return
  }
  state.currentSystem = { icon: obj.icon, name: obj.objectName }
  state.params = { topKey: obj.key, key: '', objectName: '' }
  getList()
  systemFormState.visible = false
}

const onSearch = () => {
  console.log(searchForm.key, searchForm.objectName)
  if (!searchForm.key && !searchForm.objectName) {
    respState.list = srcList
    state.expandedRowKeys = getAllChildrenKeys(srcList, [])
    return
  }
  const { matches, expandedRowKeys } = serchOnSrcList(srcList, searchForm.key, searchForm.objectName, [], [], [])
  respState.list = []
  if (!srcList) {
    state.expandedRowKeys = []
    return
  }
  const respList: Array<Obj> = []
  for (const child of srcList) {
    if (matches.includes(child.uuid)) {
      const { uuid, objectName, identifier, key, sort, type, subType, extra, icon, status, pUUID } = child
      const item = { uuid, objectName, identifier, key, sort, type, subType, extra, icon, status, pUUID }
      respList.push(item)
      getNewTree(matches, item, child.children)
    }
  }
  respState.list = respList
  state.expandedRowKeys = expandedRowKeys
}

const serchOnSrcList = (objectList: Array<Obj>, key: string, objectName: string, parents: Array<string>, matches: Array<string>, expandedRowKeys: Array<string>) => {
  for (const child of objectList) {
    let keyMatched = true
    if (key) {
      keyMatched = child.key.toLowerCase().includes(key.toLowerCase())
    }

    let nameMatched = true
    if (objectName) {
      nameMatched = child.objectName.toLowerCase().includes(objectName.toLowerCase())
    }
    if (keyMatched && nameMatched) {
      if (!matches.includes(child.uuid)) {
        matches.push(child.uuid)
      }
      parents.forEach(uuid => {
        if (!expandedRowKeys.includes(uuid)) {
          expandedRowKeys.push(uuid)
        }
        if (!matches.includes(uuid)) {
          matches.push(uuid)
        }
      })

      if (child.children && child.children.length > 0) {
        // expandedRowKeys.push(child.uuid)
        // 如果父级匹配到无需展示，可以删掉下面两行代码
        const children = getAllChildrenKeys(child.children, [])
        matches.push(...children)
      }
    }
    const newParents = [...parents, child.uuid]
    if (child.children && child.children.length > 0) {
      serchOnSrcList(child.children, key, objectName, newParents, matches, expandedRowKeys)
    }
  }

  return { matches, expandedRowKeys }
}

const getAllChildrenKeys = (objectList: Array<Obj>, keys: Array<string>): Array<string> => {
  for (const child of objectList) {
    keys.push(child.uuid)
    if (child.children && child.children.length > 0) {
      getAllChildrenKeys(child.children, keys)
    }
  }
  return keys
}

const getNewTree = (matches: Array<string>, parents: Obj, objectList?: Array<Obj>) => {
  if (!objectList) {
    return
  }
  for (const child of objectList) {
    if (matches.includes(child.uuid)) {
      const { uuid, objectName, identifier, key, sort, type, subType, extra, icon, status, pUUID } = child
      const item = { uuid, objectName, identifier, key, sort, type, subType, extra, icon, status, pUUID }
      if (!parents.children) {
        parents.children = [item]
      } else {
        parents.children.push(item)
      }
      getNewTree(matches, item, child.children)
    }
  }
}

let exportData: Array<Array<any>> = []

const onExport = () => {
  let content = `即将导出当系统"${state.currentSystem.name}"所有的菜单和操作`
  if (searchForm.key !== '' || searchForm.objectName !== '') {
    content = '即将导出当前系统筛选的菜单和操作'
  }

  Modal.confirm({
    title: '批量导出菜单和操作',
    icon: createVNode(ExclamationCircleOutlined),
    content,
    okText: '确认',
    cancelText: '取消',
    onOk() {
      const DataTitles = ['uuid', '父级uuid', '菜单名称/操作名称', '类型', '菜单类型/请求方式', '菜单路径/操作uri', '状态', '排序', '图标', '标识符']
      exportData = [DataTitles]
      getExportArr(respState.list)
      const ws = XLSX.utils.aoa_to_sheet(exportData)
      ws['!cols'] = [{ wpx: 120 }, { wpx: 120 }, { wpx: 180 }, { wpx: 60 }, { wpx: 150 }, { wpx: 200 }, { wpx: 60 }, { wpx: 90 }, { wpx: 120 }, { wpx: 120 } ]
      const wb = XLSX.utils.book_new()
      XLSX.utils.book_append_sheet(wb, ws, 'Sheet1')
      XLSX.writeFile(wb, `${state.currentSystem.name}菜单和操作.xlsx`)
    }
  })
}

const getExportArr = (arr: Array<Obj>) => {
  arr.forEach(item => {
    let typ = ''
    let subType = ''
    let status = '启用'
    if (item.status === 0) {
      status = '停用'
    }
    if (item.type === 2) {
      typ = '菜单'
      if (item.subType === 1) {
        subType = '菜单'
      } else if (item.subType === 2) {
        subType = '菜单组'
      } else {
        subType = '隐藏菜单'
      }
    } else {
      typ = '操作'
      if (item.subType === 1) {
        subType = 'GET'
      } else if (item.subType === 2) {
        subType = 'POST'
      } else if (item.subType === 3) {
        subType = 'PUT'
      } else if (item.subType === 4) {
        subType = 'PATCH'
      } else {
        subType = 'DELETE'
      }
    }
    const dataItem = [item.uuid, item.pUUID, item.objectName, typ, subType, item.key, status, item.sort, item.icon, item.identifier]
    exportData.push(dataItem)
    if (item.children && item.children.length > 0) {
      getExportArr(item.children)
    }
  })
}

const importLogContentRef = ref()

const onImport = () => {
  Modal.confirm({
    title: '批量导入菜单和操作',
    icon: createVNode(ExclamationCircleOutlined),
    content: `即将导入到系统“${state.currentSystem.name}”, 确定吗？`,
    okText: '确认',
    cancelText: '取消',
    onOk() {
      let input = document.createElement('input')
      input.type = 'file'
      input.accept = 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
      input.addEventListener('change', doImportFileParse, false)
      input.click()
    }
  })
}

const doImportFileParse = (evt: any) => {
  try {
    const file = evt.target.files[0]
    const reader = new FileReader()
    reader.onload = function (e: any) {
      const workbook = XLSX.read(e.target.result, { type: 'binary' })
      const wsname = workbook.SheetNames[0]// 取第个Sheet
      const ws = XLSX.utils.sheet_to_json(workbook.Sheets[wsname])// 生成纯数据
      const importData: Array<Obj> = []
      if (ws && ws.length) {
        for (let i = 0; i < ws.length; i++) {
          let { uuid, '父级uuid': pUUID, '菜单名称/操作名称': objectName, '类型': type, '菜单类型/请求方式': subType, '菜单路径/操作uri': key, '状态': status, '排序': sort, '图标': icon, '标识符': identifier } = ws[i] as any
          if (uuid == '') {
            message.error(`第${i+2}行,uuid列是空的`)
            return
          }
          if (key == '') {
            message.error(`第${i+2}行(uuid是${uuid}),菜单路径/操作uri是空的`)
            return
          }
          let statusVal = -1
          if (status === '启用') {
            statusVal = 1
          } else if (status === '停用') {
            statusVal = 0
          } else {
            message.error(`第${i+2}行(uuid是${uuid}),状态的值异常`)
            return
          }
          let typeVal = -1
          let subTypeVal = -1
          if (type === '菜单') {
            typeVal = 2
            if (subType === '菜单') {
              subTypeVal = 1
            } else if (subType === '菜单组') {
              subTypeVal = 2
            } else if (subType === '隐藏菜单')  {
              subTypeVal = 3
            } else {
              message.error(`第${i+2}行,菜单类型/请求方式的值异常`)
              return
            }
          } else if (type === '操作') {
            typeVal = 3
            if (subType === 'GET') {
              subTypeVal = 1
            } else if (subType === 'POST') {
              subTypeVal = 2
            } else if (subType === 'PUT') {
              subTypeVal = 3
            } else if (subType === 'PATCH') {
              subTypeVal = 4
              subType = 'PATCH'
            } else if (subType === 'DELETE') {
              subTypeVal = 5
            } else {
              message.error(`第${i+2}行,菜单类型/请求方式的值异常`)
              return
            }
          } else {
            message.error(`第${i+2}行,类型的值异常`)
            return
          }
          importData.push({ uuid, pUUID, objectName, type: typeVal, subType: subTypeVal, key, status: statusVal, sort, icon, identifier })
        }
      }
      if (importData.length > 0) {
        doImmport(importData)
      } else {
        message.error('导入数据异常')
      }
    }
    reader.readAsBinaryString(file)
  } catch {
    message.error('导入数据文件发生错误')
  }
}

interface ImportItem {
  uuid: string
  objectName: string
  status: string
  msg: string
}

interface ImportState {
  visible: boolean
  percent: number
  total: number // 共需要导入的数量
  timeCount: number
  data: Array<ImportItem>
  statistics: Record<'success' | 'ignore' | 'failed', number>
}

const importState = reactive<ImportState>({
  visible: false,
  percent: 0,
  total: 0,
  timeCount: 0,
  statistics: {
    success: 0,
    ignore: 0,
    failed: 0
  },
  data: []
})

let timer: NodeJS.Timer

const doImmport = async (importData: Array<any>) => {
  importState.visible = true
  // 打开导入弹窗，展示导入进度
  importState.statistics = {
    success: 0,
    ignore: 0,
    failed: 0
  }
  importState.total = importData.length
  importState.data =[]
  importState.percent = 0
  importState.timeCount = 0
  if (timer) {
    clearInterval(timer)
  }
  timer = setInterval(() => {
    importState.timeCount += 0.1
  }, 100)
  for (let i = 0; i < importData.length; i++) {
    const item = importData[i]
    item.topKey = state.params.topKey
    console.log('import item', item)
    importState.data.push({ uuid: item.uuid, objectName: item.objectName, status: 'doing', msg: '' })
    try {
      const { status, msg } = await importObject(item)
      importState.data[importState.data.length -1].status = status
      importState.data[importState.data.length -1].msg = msg
      if (status === 'success') {
        importState.statistics.success++
      } else if(status === 'ignore') {
        importState.statistics.ignore++
      } else if(status === 'failed') {
        importState.statistics.failed++
      }
    } catch (err: any) {
      importState.data[importState.data.length -1].status = 'failed'
      importState.data[importState.data.length -1].msg = err.message
      importState.statistics.failed++
    }
    importState.percent = Math.ceil(importState.data.length / importData.length * 100)
    if (importState.percent == 100) {
      clearInterval(timer)
      getList()
    }
    nextTick(() => {
      importLogContentRef.value.scrollTop = importLogContentRef.value.scrollHeight
    })
  }
}

const onImportCancel = () => {
  clearInterval(timer)
}

onBeforeUnmount(() => {
  clearInterval(timer)
})
</script>
<style scoped>
.import-log{
  overflow: hidden;
}
.import-log__content{
  color: rgb(138, 138,138);
  margin: 10px 20px;
  background-color: black;
  padding: 10px;
  height: 200px;
  overflow-y: scroll;
  font-family: 'Courier New', Courier, monospace;
  font-size: 12px;
}
.import-ready{
  color: yellow
}
.import-doing{
  color: #2db7f5
}
.import-success{
  color: green;
}
.import-failed{
  color: red;
}
.import-ignore{
  color: rgb(138, 138,138);
}
.log__content_statistics{
  margin: 0 20px 20px;
}
</style>
