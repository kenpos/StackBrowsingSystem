<template>
  <div class="hello">
    <button @click="sendRequest">リクエスト送信</button>
    <h1>取得結果</h1>
    <p>GET:<br/>{{hello}}</p>

    <button @click="sendRequestEdinet">有価証券報告書取得</button>
    <h1>取得結果</h1>
    <table id="metadata"> 
      <tr>
          <th>KEY</th>
          <th>DocID</th>
          <th>edinetCode</th>
          <th>filerName</th>
          <th>docDescription</th>
          <th>XbrlFlag</th>
          <th>PdfFlag</th>
      </tr>
      <tr v-for="(value, key) in edinet.results" v-bind:key="value.seqNumber">
          <td>{{ key }}</td>
          <td>{{ value.DocID }}</td>
          <td>{{ value.edinetCode }}</td>
          <td>{{ value.filerName }}</td>
          <td>{{ value.docDescription }}</td>
          <td>{{ value.XbrlFlag }}</td>
          <td>{{ value.PdfFlag }}</td>
      </tr>
  </table>
  
    <p>GET:<br/>{{edinet}}</p>
  </div>
</template>

<script>
import axios from "axios"

export default {
  name: 'App',
  data:()=>{
    return{
      edinet:"",
      hello:"",
    }
  },
  methods: {
    sendRequest: async function(){
      const getRequestNoParam = await axios.get("/hello")
      this.hello = getRequestNoParam.data
    },
    sendRequestEdinet: async function(){
      const getRequestNoParam = await axios.get("/getEdinetAPI")
      this.edinet = getRequestNoParam.data
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
