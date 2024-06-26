<template>
  <div>
    <div>
      <a-button type="primary" @click="visible = true">新增标签</a-button>
      <a-modal
        v-model:open="visible"
        title="新增标签"
        ok-text="提交"
        cancel-text="取消"
        @ok="addTag"
      >
        <a-form ref="formRef" :model="formState" layout="vertical" name="form_in_modal">
          <a-form-item
            name="name"
            label="名称"
            :rules="[{ required: true, message: '请输入标签名称' }]"
          >
            <a-input v-model:value="formState.name" />
          </a-form-item>
          <a-form-item
            name="route"
            label="前端路由"
            :rules="[{ required: true, message: '请输入前端路由' }]"
          >
            <a-input v-model:value="formState.route" />
          </a-form-item>
          <a-form-item
            label="是否启用"
            name="enabled"
            class="collection-create-form_last-form-item"
          >
            <a-radio-group v-model:value="formState.enabled">
              <a-radio :value="true">true</a-radio>
              <a-radio :value="false">false</a-radio>
            </a-radio-group>
          </a-form-item>
        </a-form>
      </a-modal>
    </div>
    <div>
      <a-table :columns="columns" :data-source="data" :pagination="pagination" @change="change">
        <template #bodyCell="{ column, text, record }">
          <template v-if="column.dataIndex === 'create_time'">
            {{ dayjs.unix(text).format('YYYY-MM-DD HH:mm:ss') }}
          </template>

          <template v-if="column.dataIndex === 'update_time'">
            {{ dayjs.unix(text).format('YYYY-MM-DD HH:mm:ss') }}
          </template>

          <template v-if="column.key === 'enabled'">
            <a-switch v-model:checked="record.enabled" @change="changeTagEnabled(record)" />
          </template>

          <template v-else-if="column.dataIndex === 'operation'">
            <a-popconfirm v-if="data.length" title="确认删除？" @confirm="deleteTag(record.id)">
              <a>删除</a>
            </a-popconfirm>
          </template>
        </template>
      </a-table>
    </div>
  </div>
</template>
<script lang="ts" setup>
import originalAxios from 'axios'
import { ref, reactive, computed } from 'vue'
import type { FormInstance } from 'ant-design-vue'
import type { PageRequest } from '@/interfaces/Common'
import { message } from 'ant-design-vue'
import dayjs from 'dayjs'
import {
  AddTag,
  ChangeTagEnabled,
  DeleteTag,
  GetTag,
  type Tag,
  type TagRequest
} from '@/interfaces/Tag'

document.title = '标签列表 - 后台管理'

const columns = [
  {
    title: '名称',
    dataIndex: 'name',
    key: 'name'
  },
  {
    title: '路由',
    dataIndex: 'route',
    key: 'route'
  },
  {
    title: '文章数量',
    key: 'post_count',
    dataIndex: 'post_count'
  },
  {
    title: '状态',
    key: 'enabled',
    dataIndex: 'enabled'
  },
  {
    title: '创建时间',
    key: 'create_time',
    dataIndex: 'create_time'
  },
  {
    title: '最后一次修改的时间',
    key: 'update_time',
    dataIndex: 'update_time'
  },
  {
    title: 'operation',
    dataIndex: 'operation'
  }
]

const data = ref<Tag[]>([])

const pageReq = ref<PageRequest>({
  pageNo: 1,
  pageSize: 5,
  sortField: 'create_time',
  sortOrder: 'desc'
} as PageRequest)

const total = ref(0)

const pagination = computed(() => ({
  total: total.value,
  current: pageReq.value.pageNo,
  pageSize: pageReq.value.pageSize
}))

const getTags = async () => {
  try {
    const response: any = await GetTag(pageReq.value)
    data.value = response.data.data?.list || []
    total.value = response.data.data?.totalCount || 0
  } catch (error) {
    console.log(error)
  }
}

getTags()

// 添加标签
const formRef = ref<FormInstance>()
const visible = ref(false)
const formState = reactive<TagRequest>({
  name: '',
  route: '',
  enabled: true
})

const addTag = () => {
  if (formRef.value) {
    formRef.value
      .validateFields()
      .then(async () => {
        try {
          const response: any = await AddTag(formState)
          if (response.data.code !== 0) {
            message.error(response.data.message)
            return
          }
          message.success('添加成功')
          visible.value = false
          if (formRef.value) {
            formRef.value.resetFields()
          }
          await getTags()
        } catch (error) {
          if (originalAxios.isAxiosError(error)) {
            // 这是一个由 axios 抛出的错误
            if (error.response) {
              if (error.response.status === 409) {
                message.error('标签名称或路由重复')
                return
              }
            } else if (error.request) {
              // 请求已发出，但没有收到响应
              console.log('No response received:', error.request)
            } else {
              // 在设置请求时触发了一个错误
              console.log('Error Message:', error.message)
            }
          }
          message.error('添加失败')
        }
      })
      .catch((info) => {
        console.log('Validate Failed:', info)
        message.warning('请检查表单是否填写正确')
      })
  }
}

const changeTagEnabled = async (record: Tag) => {
  try {
    const response: any = await ChangeTagEnabled(record.id, record.enabled)
    if (response.data.code !== 0) {
      message.error(response.data.message)
      return
    }
    message.success('修改成功')
  } catch (error) {
    console.log(error)
  }
  await getTags()
}

// 删除
const deleteTag = async (id: string) => {
  try {
    const response: any = await DeleteTag(id)
    if (response.data.code !== 0) {
      message.error(response.data.message)
      return
    }
    message.success('删除成功')
    await getTags()
  } catch (error) {
    console.log(error)
    if (originalAxios.isAxiosError(error)) {
      // 这是一个由 axios 抛出的错误
      if (error.response) {
        if (error.response.data.status === 404) {
          message.error('id 不存在')
          return
        }
      } else if (error.request) {
        // 请求已发出，但没有收到响应
        console.log('No response received:', error.request)
      } else {
        // 在设置请求时触发了一个错误
        console.log('Error Message:', error.message)
      }
    }
    message.error('删除失败')
  }
}

const change = (pg: any) => {
  pageReq.value.pageNo = pg.current
  pageReq.value.pageSize = pg.pageSize
  getTags()
}
</script>

<style scoped>
.collection-create-form_last-form-item {
  margin-bottom: 0;
}
</style>
