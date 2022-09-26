<template>
  <div class="M">
    <div class="top">
      <h2>海绵摆摆电影院2021年电影播放量TOP8</h2>
    </div>
    <div id="main" style="width:700px;height:600px"></div>
  </div>
</template>

<script>
export default {
  mounted () {
    var chartDom = document.getElementById('main');
var myChart = this.$echarts.init(chartDom);
var option;

option = {
  legend: {
    top: 'bottom'
  },
  toolbox: {
    show: true,
    feature: {
      mark: { show: true },
      dataView: { show: true, readOnly: false },
      restore: { show: true },
      saveAsImage: { show: true }
    }
  },
  series: [
    {
      name: '',
      type: 'pie',
      radius: [50, 250],
      center: ['50%', '50%'],
      roseType: 'area',
      itemStyle: {
        borderRadius: 10
      },
      data: [
        { value: 40, name: '心灵奇旅' },
        { value: 38, name: '花束般的恋爱' },
        { value: 32, name: '困在时间里的父亲' },
        { value: 32, name: '沙丘' },
        { value: 30, name: '失控玩家' },
        { value: 26, name: '你好,李焕英' },
        { value: 22, name: '无辜者' },
        { value: 18, name: '倒数时刻' },
      ]
    }
  ]
};

option && myChart.setOption(option);
  }
}
</script>

<style>
  .top{
    /* margin-left: 50px; */
    text-align: center;
}
  .top h2 {
    padding-bottom: 20px;
  }
  #main {
    float: left;
    margin-top: 50px;
    margin-bottom: 50px;
    margin-left: 200px;
  }
</style>