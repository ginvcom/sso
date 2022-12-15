<template>
  <template v-if="c.length > 0 && visibleChildren(c)">
   <template v-if="m.t == 1">
      <a-sub-menu :key="n" :title="m.n">
        <template #icon v-if="m.i">
          <icon-font :type="m.i" />
        </template>
        <menu-item v-for="child in c" :key="child.n" v-bind="child" />
      </a-sub-menu>
    </template>
    <a-menu-item-group :key="n" :title="m.n" v-else-if="m.t == 2">
      <menu-item v-for="child in c" :key="child.n" v-bind="child">
        <template #icon v-if="m.i">
          <icon-font :type="m.i" />
        </template>
      </menu-item>
    </a-menu-item-group>
  </template>
  <template v-else>
    <template v-if="m.t == 1">
      <a-menu-item :key="n">
        <template #icon v-if="m.i">
          <icon-font :type="m.i" />
        </template>
        <router-link :to="p">{{m.n}}</router-link>
      </a-menu-item>
    </template>
    <a-menu-item-group :key="n" :title="m.n" v-else-if="m.t == 2">
      <template #icon v-if="m.i">
        <icon-font :type="m.i" />
      </template>
    </a-menu-item-group>
  </template>
</template>
<script setup lang="ts">
import { createFromIconfontCN } from '@ant-design/icons-vue'
// import { Menu, Meta } from '@/api/auth'
// 需要用字面变量

export interface Meta {
	n: string // 菜单名
	i: string // 菜单icon
	u: string // 唯一标识符
	t: number // 菜单类型，1菜单，2菜单组, 3隐藏菜单
}

export interface Menu {
	n: string // 路由命名
	p: string // 路由路径
	m: Meta // 路由meat
	c: Array<Menu> // 路由children
}
const props = withDefaults(defineProps<Menu>(), {})

const visibleChildren = (c: Array<Menu>) => {
  let visible = false
  for (const child of c) {
    if (child.m.t !== 3) {
      visible = true
      break
    }
  }
  return visible
}

const iconFont = createFromIconfontCN({
  scriptUrl: '//at.alicdn.com/t/c/font_3601463_tk54gkmpecb.js',
})
</script>
