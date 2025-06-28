<template>
  <div id="app">
    <div v-if="false">
      <ElButton size="mini" @click="addAll">添加全部</ElButton>
      <ElButton size="mini" @click="cancelAll">取消全部</ElButton>
    </div>
    <el-checkbox-group v-model="selected" :style="{ lineHeight: '22px' }">
      <el-checkbox
        @change="handleChange($event, key)"
        v-for="(value, key) in data"
        :key="key"
        :label="key"
        :name="key"
        :disabled="false"
        :border="false"
        >{{ key.replace(/^custommods\//, "") }}</el-checkbox
      >
    </el-checkbox-group>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      data: {},
      selected: [],
    };
  },
  methods: {
    async handleChange(event, item) {
      var result;
      if (event == true) {
        const elLoadingComponent = this.$loading();
        result = await axios({
          method: "POST",
          url: "http://127.0.0.1:9090/mod",
          data: {
            type: "add",
            dirPath: item,
          },
        }).finally(() => {
          elLoadingComponent.close();
        });
      }
      if (event == false) {
        const elLoadingComponent = this.$loading();
        result = await axios({
          method: "POST",
          url: "http://127.0.0.1:9090/mod",
          data: {
            type: "remove",
            dirPath: item,
          },
        }).finally(() => {
          elLoadingComponent.close();
        });
      }

      const { status, errors = [] } = result.data || {};
      if (status != 200) {
        this.$message.error(`${errors}`);
        this.init();
      }
    },
    init() {
      const elLoadingComponent = this.$loading();
      axios({
        method: "GET",
        url: "http://127.0.0.1:9090/mods",
      })
        .then((result) => {
          this.data = result.data;
          this.selected = Object.entries(this.data)
            .filter((item) => {
              return item[1] == true;
            })
            .map((item) => {
              return item[0];
            });
        })
        .finally(() => {
          elLoadingComponent.close();
        });
    },
    addAll() {
      // const keys = Object.entries(this.data).filter(item=>{
      //     return item[1] == false;
      //   }).map(item=>{
      //     return item[0]
      //   })
    },
    cancelAll() {
      //  const keys = Object.entries(this.data).filter(item=>{
      //   return item[1] == true;
      // }).map(item=>{
      //   return item[0]
      // })
    },
  },
  created() {
    this.init();
  },
};
</script>