<template>
  <div class="table-box">
    <!-- 标题 -->
     <div class="title">
      <h2>最简单的crud demo</h2>
     </div>
     
     <!-- query -->
     <div class="query-box">
      <el-input-tag class="query-input" v-model="queryInput" placeholder="请输入名字搜索🔍" @change="handleQueryName"/>
      <div class="btn-list">
      <el-button type="primary" @click="handleAdd">增加</el-button>
      <el-button type="danger" @click="handleDelList">删除多选</el-button>
      </div>
      
     </div> 

     <!-- table -->
    <el-table 
      border
      ref="multipleTableRef"
      :data="tableData"
      style="width: 100%"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column prop="name" label="姓名" width="120" />
      <el-table-column prop="email" label="邮箱" width="120" />
      <el-table-column prop="phone" label="电话" width="120" />
      <el-table-column prop="status" label="状态" width="120" />
      <el-table-column prop="address" label="地址" width="600" />

      <el-table-column fixed="right" label="操作" min-width="120">
        <template #default="scope">
          <el-button link type="primary" size="small" @click="handleRowDel(scope.row)" style="color: #F56C6C;">删除</el-button>
          
          <el-button link type="primary" size="small" @click="handleEdit(scope.row)">编辑</el-button>
        </template>
      </el-table-column>
  </el-table>



    <!-- dialog -->
  <el-dialog v-model="dialogFormVisible" :title="dialogType === 'add'?'新增':'编辑' "width="500">
    <el-form :model="tableForm">
      <el-form-item label="姓名" :label-width="80">
        <el-input v-model="tableForm.name" autocomplete="off" />
      </el-form-item>
      <el-form-item label="邮箱" :label-width="80">
        <el-input v-model="tableForm.email" autocomplete="off" />
      </el-form-item>
      <el-form-item label="电话" :label-width="80">
        <el-input v-model="tableForm.phone" autocomplete="off" />
      </el-form-item>
      <el-form-item label="状态" :label-width="80">
        <el-input v-model="tableForm.status" autocomplete="off" />
      </el-form-item>
      <el-form-item label="地址" :label-width="80">
        <el-input v-model="tableForm.address" autocomplete="off" />
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button type="primary" @click="dialogConfirm">确认</el-button>
      </div>
    </template>
  </el-dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import request from './utils/request'
  /** 数据 */
  let queryInput = ref('')
  let tableData = ref([
  {
    id: "1",
    name: 'Tom1',
    email:'132@qq.com',
    phone: '13212345678',
    status: '在职',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    id: "2",
    name: 'Tom2',
    email:'132@qq.com',
    phone: '13212345678',
    status: '在职',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    id: "3",
    name: 'Tom3',
    email:'132@qq.com',
    phone: '13212345678',
    status: '在职',
    address: 'No. 189, Grove St, Los Angeles',
  },
])
  let tableDataCopy = ref([])
  // 初始化拷贝原始数据
  tableDataCopy.value = [...tableData.value]
  let multipleSelection = ref([])
  let dialogFormVisible = ref(false)
  let tableForm = ref({
    ID: '',
    name: '',
    email: '',
    phone: '',
    status: '',
    address: '',
  })
  let dialogType = ref('add')

  /* 方法 */ 

const getTableData = async () =>{
  let res = await request.get('/list')
  console.log(res);
  tableData.value = res
  
}
getTableData()
//搜索
const handleQueryName = async (val)=>{
  // if (val.length>0) {
  //   tableData.value = tableData.value.filter(item => (item.name).toLowerCase().match(val.toLowerCase()))
  // }else{
  //   tableData.value = tableDataCopy.value
  // }

  let res = await request.get(`/list/${val}`)
  tableData.value = res
}
//编辑
const handleEdit = (row)=>{
  dialogFormVisible.value = true
  dialogType.value = 'edit'
  tableForm.value = {...row}
}

 // 删除单条数据 
const handleRowDel = async({ID}) =>{
  // //1.通过id获取到条目对应的索引
  // let index = tableData.value.findIndex(item => item.id === id)
  // //2.通过索引删除对应条目数据
  // tableData.value.splice(index, 1)

  await request.delete(`/delete/${ID}`)
  await getTableData()

}
// 删除多条数据
const handleDelList = ()=>{
  multipleSelection.value.forEach(ID => {
    handleRowDel({ID})
  })
  multipleSelection.value = []
}
//选中
const handleSelectionChange = (val) => {
  // multipleSelection.value = val
  multipleSelection.value = []
  val.forEach(item => {
    multipleSelection.value.push(item.ID)
  })
}

//新增
const handleAdd = () => {
  dialogFormVisible.value = true
  tableForm.value = {}
  dialogType.value = 'add'
}

//确认
const dialogConfirm = async() => {
  dialogFormVisible.value = false

  //判断是新增还是更改
  if (dialogType.value === 'add') {
    // 1.拿到数据
    // 2.添加到table
    // tableData.value.push(
    //   {
    //     id:(tableData.value.length + 1).toString(),
    //     ...tableForm.value
    //   }
    // )

    //添加数据
    await request.post('/add',{
      ...tableForm.value
    })
    //刷新数据
    await getTableData()

  }else{
    // // 1.获取到当前的这条索引
    // let index = tableData.value.findIndex(item => item.id === tableForm.value.id) 
    // // 2.替换当前索引值对应的数据
    // tableData.value[index] = tableForm.value

    //更改数据
    await request.put(`/update/${tableForm.value.ID}`,{
      ...tableForm.value
    })
    await getTableData()
  }

  
}

</script>

<style scoped>
.table-box { 
  margin: 300px auto;
  width: 800px;
}
.title {
  text-align: center;
}
.query-box {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
  /* display: flex将容器设为弹性盒子，其子元素自动成为弹性项目 
     justify-content 控制主轴（默认水平方向）的对齐方式
     space-between 适用于导航栏、卡片布局等需要两端对齐的场景*/
}
.query-input {
  width: 300px;
}
</style>
