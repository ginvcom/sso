<template>
  <div class="sign-in">
    <div class="sign-in__left">
      <div class="sign-in__logo">
        <a href="/" class="sign-in__link">
          <div class="logo">
            <div class="logo-inner">
              <div class="round"></div>
              <div class="line">
              </div>
              <div class="line line2">
              </div>
            </div>
          </div>
          <p>
            Single Sign On
          </p>
        </a>
      </div>
      <div class="center-on-page">
        <div class="logo">
          <div class="logo-inner">
            <div class="round"></div>
            <div class="line">
            </div>
            <div class="line line2">
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="sign-in__right">
      <div class="sign-in__form">
        <h2 class="sign-in__title">Welcome Back ✨</h2>
        <p class="sign-in__subtitle">Please sign in to continue</p>
        <form
          class="ant-form ant-form-vertical"
          :action="state.action"
          enctype="application/x-www-form-urlencoded"
          method="post"
        >
          <a-form-item
            label="Mobile"
            name="mobile"
            :rules="[{ required: true, message: 'Please input your username!' }]"
          >
            <a-input name="mobile" v-model:value="formState.mobile" />
          </a-form-item>

          <a-form-item
            label="Password"
            name="password"
            :rules="[{ required: true, message: 'Please input your password!' }]"
          >
            <a-input-password name="password" v-model:value="formState.password" />
          </a-form-item>
          <div class="sign-in__extra">
          <a-form-item name="remember">
            <a-checkbox name="remember" v-model:checked="formState.remember">Remember me</a-checkbox>
          </a-form-item>
          <!-- <router-link to="/">forgot password?</router-link> -->
          </div>
          <a-form-item>
            <a-button type="primary" block html-type="submit">Sign In</a-button>
          </a-form-item>
        </form>
        <p class="sign-in__time">{{state.year}}年 <a-divider type="vertical" /> 第{{state.weekNumber}}周</p>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { reactive, onBeforeMount } from 'vue'
import { SignInReq } from './api/auth'
import { useRoute } from 'vue-router'

interface State {
  action: string
  year: number | null
  weekNumber: number | null
}

const state = reactive<State>({
  action: '/sign-in',
  year: null,
  weekNumber: null
})

const route = useRoute()

onBeforeMount(() => {
  const baseRUL = import.meta.env.VITE_BASEURL
  if (baseRUL) {
    state.action = baseRUL + '/sign-in'
  }
  const serviceCode = route.query.serviceCode
  if (serviceCode) {
    state.action = state.action + '?serviceCode=' + serviceCode
  }
  getWeekNumber()
})

const formState = reactive<SignInReq>({
  mobile: '',
  password: '',
  remember: true
})

// const onFinish = (values: SignInReq) => {
//   signIn(values).then(res => {
//     console.log(res)
//     if (res) {
//       setCookie('user', JSON.stringify(res))
//     }
//   })
// }

const getWeekNumber = () => {
  /**
  * a = d = 当前日期
  * b = 6 - w = 当前周的还有几天过完(不算今天)
  * a + b 的和在除以7 就是当天是当前月份的第几周
  */
  var date = new Date()
  state.year = date.getFullYear()
  const firstDay = new Date(state.year + '-01-01')
  let w = firstDay.getDay()
  let d = date.getDate()
  if (w == 0) {
    w =  7
  }
  const days = Math.floor((date.valueOf() - firstDay.valueOf()) / (3600 * 1000 * 24))
  console.log(days)
  let n = Math.ceil((days - 7 + w) / 7)
  if (w > 0) {
    n++
  }
  state.weekNumber = n
}
</script>
<style scoped>
.sign-in__logo .logo {
  float: left;
  overflow: hidden;
}
.sign-in{
  display: flex;
  height: 100%;
}
.sign-in__left{
  flex: 1;
  background-color: #f8fafc;
  display: flex;
  position: relative;
}

.sign-in__logo{
  position: fixed;
}

.sign-in__logo img {
  width: 48px;
}
.sign-in__logo p{
  font-weight: bold;
  color: #334155;
  font-size: 100px;
  line-height: 200px;
  margin: 0 0 0 20px;
}
.sign-in__link{
  display: flex;
  align-items: center;
  padding: 80px;
  transform: scale(0.2);
  transform-origin: 0 0;
}
.sign-in__right{
  width: 500px;
  background-color: #fff;
  text-align: center;
  display: flex;
  align-items: center;
  justify-content: center;
}

.sign-in__form{
  width: 400px;
}
h2.sign-in__title{
  color: #475569;
  font-size: 1.5rem;
  line-height: 2rem;
}
.sign-in__subtitle{
  color: #94a3b8;
}

.sign-in__time{
  color: #94a3b8;
  margin-top: 80px;
  font-size: 14px;
}
.sign-in__extra{
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.sign-in__extra a{
  margin-bottom: 24px;
}
</style>
