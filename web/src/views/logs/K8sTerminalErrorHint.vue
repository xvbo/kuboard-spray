<i18n>
en:
  "referto": "If you are using Reverse Proxy, please refer to"
  "commonProblems": "Recommended Reverse Proxy configuration."
  "confirm": "Do you confirm to retry ?"
zh:
  "referto": "如果您设置了反向代理，请参考"
  "commonProblems": "Kuboard 推荐的反向代理配置。"
  "confirm": "确认重试？"
</i18n>

<template>
  <div style="text-align: left;">
    <el-dialog :title="$t('msg.prompt')" v-model="dialogVisible" width="60%" :before-close="handleClose">
      <p>{{$t('referto')}}</p>
      <!-- <KuboardHelpLink hideOnDisabled heighlight href="https://kuboard.cn/install/v3/proxy.html" :label="$t('commonProblems')"></KuboardHelpLink> -->
      <!-- <span>{{message}}</span> -->
    </el-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      dialogVisible: false,
      message: undefined
    }
  },
  methods: {
    show(message) {
      this.message = message
      this.dialogVisible = true
    },
    handleClose() {
      this.$confirm(this.$t('confirm'))
        .then(() => {
          window.location.reload()
        })
        .catch(() => { });
    }
  }
}
</script>

<style>
</style>
