/**
 * axios 二次封装
 * 
 * 全局配置
 * 
 * 响应拦截
 * 
 * request 请求的方法
 * 封装成对象调用的方式
 * request.get(url,{params:{}})
 */

import axios from 'axios'
import { ElMessage } from 'element-plus'

//全局配置
const service = axios.create({
    baseURL: "/user",//根路径
    timeout: 8000 //请求超时时间
})


//响应拦截
service.interceptors.response.use(res => {
    const { code, data, msg } = res.data

    if (code === 200) {
        //请求成功
        ElMessage.success(msg)
        return data
    } else if (code === 400) {
        //请求失败
        ElMessage.error(msg)

    }

})
// service.interceptors.response.use(
//     (res) => {
//         const { code, data, msg } = res.data

//         if (code === 200) {
//             msg && ElMessage.success(msg)
//             return data
//         } else {
//             msg && ElMessage.error(msg)
//             return Promise.reject(new Error(msg || `Error code: ${code}`))
//         }
//     },
//     (error) => {
//         const msg = error.response?.data?.msg || error.message || '请求失败'
//         ElMessage.error(msg)
//         return Promise.reject(error)
//     }
// )

//request 请求的方法
// 封装请求：统一处理get请求参数格式，调用service发送请求
function request(options) {
    // 默认请求方法为get
    options.method = options.method || 'get'

    // get请求：将data参数映射到params（适配service的参数约定）
    if (options.method.toLowerCase() == 'get') {
        options.params = options.data
    }
    //执行底层请求并返回Promise
    return service(options)
}

//request.get(url,{params:{}})
// 第一步：定义要封装的请求方法列表（常用的4种HTTP方法）
['get', 'post', 'delete', 'put'].forEach(item => {
    // 第二步：为request挂载快捷方法（比如request.get、request.post）
    request[item] = (url, data) => {
        // 第三步：调用原来的request函数，自动组装配置
        return request({
            url,    // 把传入的url传给原request的url
            data,   // 把传入的data传给原request的data
            method: item  // 自动把method设为当前循环的方法（get/post等）
        })
    }
})

export default request
