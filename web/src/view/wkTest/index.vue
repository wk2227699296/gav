<template>
  <div class="wk-test">
    <div class="title">
      <h1>wkTest</h1>
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
        <el-table-column label="学生" prop="student">
          <template #default="{ row }">
            <span
              v-for="item in row.student"
              :key="item.ID"
              style="margin-right: 4px"
              >{{ item.name }}</span
            >
          </template>
        </el-table-column>
        <el-table-column label="class" prop="classList">
          <template #default="{ row }">
            <span
              v-for="item in row.classList"
              :key="item.ID"
              style="margin-right: 4px"
              >{{ item.name }}</span
            >
          </template>
        </el-table-column>
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
        <el-form-item label="学生" prop="student">
          <el-select
            v-model="wkInfo.student"
            placeholder=""
            allow-create
            multiple
            value-key="ID"
            @focus="selectStudentFocus"
          >
            <el-option
              v-for="item in wkStudentList"
              :key="item.ID"
              :label="item.name"
              :value="item"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="class" prop="class">
          <el-select
            v-model="wkInfo.class"
            placeholder=""
            allow-create
            multiple
            @focus="selectClassFocus"
          >
            <el-option
              v-for="item in wkClassList"
              :key="item.ID"
              :label="item.name"
              :value="item.ID"
            ></el-option>
          </el-select>
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
  getWkInfoList,
  createWkInfo,
  updateWkInfo,
  deleteWkInfoByIds,
} from "@/api/wkInfo";
import { getWkStudentList } from "@/api/wkStudent";
import { getWkClassList } from "@/api/wkClass";
import { ref } from "vue";
import _ from "lodash";
const wkInfoList = ref([]);
const wkStudentList = ref([]);
const wkClassList = ref([]);
const wkInfo = ref({
  name: "",
  age: "",
});

const addOrUpdate = ref(false);
const dialogVisible = ref(false);

const getWkInfo = async () => {
  const res = await getWkInfoList();
  res.data.list.forEach((item) => {
    item.class = JSON.parse(item.class);
  });
  wkInfoList.value = res.data.list;
};

const openDialog = (type, row) => {
  console.log(type, row, "type");
  addOrUpdate.value = type === "add";
  dialogVisible.value = true;
  row && (wkInfo.value = _.cloneDeep(row));
};

const addWkInfo = async () => {
  console.log(wkInfo.value, "wkInfo");
  wkInfo.value.class = JSON.stringify(wkInfo.value.class);
  if (addOrUpdate.value) {
    // 新增
    await createWkInfo(wkInfo.value);
  } else {
    // 修改

    await updateWkInfo(wkInfo.value);
  }

  dialogVisible.value = false;
  getWkInfo();
  wkInfo.value = {
    name: "",
    age: "",
  };
};

const deleteWkInfo = async (id) => {
  console.log(id, "id");
  await deleteWkInfoByIds({ ids: [id] });
  getWkInfo();
};

const selectStudentFocus = async () => {
  const res = await getWkStudentList();
  wkStudentList.value = res.data.list;
};

const selectClassFocus = async () => {
  const res = await getWkClassList();
  wkClassList.value = res.data.list;
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
