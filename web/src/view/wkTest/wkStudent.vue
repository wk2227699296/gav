<template>
  <div class="wk-test">
    <div class="title">
      <h1>wkStudent</h1>
      <div>
        <el-button type="primary" @click="getWkInfo">获取数据</el-button>
        <el-button type="primary" @click="openDialog('add')"
          >新增数据</el-button
        >
      </div>
    </div>
    <div class="card">
      <el-table :data="wkInfoList" border>
        <el-table-column label="姓名" prop="name"></el-table-column>
        <el-table-column label="年龄" prop="age"></el-table-column>
        <el-table-column label="操作">
          <template #default="{ row }">
            <el-button
              type="primary"
              size="mini"
              @click="openDialog('update', row)"
            >
              编辑
            </el-button>
            <el-button type="danger" size="mini" @click="deleteWkInfo(row.ID)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog
      :title="addOrUpdate ? '新增数据' : '编辑数据'"
      v-model="dialogVisible"
    >
      <el-form :model="wkInfo" label-width="80px">
        <el-form-item
          label="姓名"
          prop="name"
          :rules="[
            {
              required: true,
              message: 'Please input Activity name',
              trigger: 'blur',
            },
          ]"
        >
          <el-input v-model="wkInfo.name"></el-input>
        </el-form-item>
        <el-form-item
          label="年龄"
          prop="age"
          :rules="[
            {
              required: true,
              message: 'Please input Activity age',
              trigger: 'blur',
            },
          ]"
        >
          <el-input v-model="wkInfo.age"></el-input>
        </el-form-item>
        <el-form-item
          label="邮箱"
          prop="email"
          :rules="[
            {
              required: true,
              message: 'Please input Activity email',
              trigger: 'blur',
            },
          ]"
        >
          <el-input v-model="wkInfo.email"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="addWkInfo">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  getWkStudentList,
  createWkStudent,
  updateWkStudent,
  deleteWkStudentByIds,
} from "@/api/wkStudent";
import { ref } from "vue";
import _ from "lodash";

const wkInfoList = ref([]);
const wkInfo = ref({
  name: "",
  age: "",
  email: "",
});

const addOrUpdate = ref(false);
const dialogVisible = ref(false);

const getWkInfo = async () => {
  const res = await getWkStudentList();
  wkInfoList.value = res.data.list;
};

const openDialog = (type, row) => {
  addOrUpdate.value = type === "add";
  dialogVisible.value = true;
  row && (wkInfo.value = _.cloneDeep(row));
};

const addWkInfo = async () => {
  if (addOrUpdate.value) {
    // 新增
    await createWkStudent(wkInfo.value);
  } else {
    // 修改

    await updateWkStudent(wkInfo.value);
  }

  dialogVisible.value = false;
  getWkInfo();
  wkInfo.value = {
    name: "",
    age: "",
    email: "",
  };
};

const deleteWkInfo = async (id) => {
  console.log(id, "id");
  await deleteWkStudentByIds({ ids: [id] });
  getWkInfo();
};

getWkInfo();
</script>

<style scoped lang="scss">
.wk-test {
  h1 {
    color: red;
  }
  .card {
    background: #fff;
  }
  .title {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
}
</style>
