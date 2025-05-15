const DefaultApi = require('./js-client/dist/api/DefaultApi.js').default;
const ApiClient = require('./js-client/dist/ApiClient.js').default;
const ResourceRequest = require('./js-client/dist/model/ResourceRequest.js').default;
const GetResource200Response = require('./js-client/dist/model/GetResource200Response.js').default;
const Pod = require('./js-client/dist/model/Pod.js').default;

const client = ApiClient.instance || new ApiClient();
client.basePath = "http://localhost:8080";

const api = new DefaultApi(client);

// 构造 ResourceRequest 请求体
const req = new ResourceRequest(
  '', // group: core 资源组为空字符串
  'v1', // version
  'Pod', // kind
  'default', // namespace
  'nginx-deployment-6d7f4f8b96-4bs7w' // name
);

api.getResource(req, (error, data, response) => {
  if (error) {
    console.error('调用出错：', error);
  } else {
    // data.data 是资源对象的 JSON 字符串
    try {
      const podObj = Pod.constructFromObject(JSON.parse(data.data));
      console.log('Pod 名称：', podObj.metadata && podObj.metadata.name);
      console.log('Pod 状态：', podObj.status && podObj.status.phase);
      console.log('完整 Pod 对象：', podObj);
    } catch (e) {
      console.error('Pod 解析失败：', e);
      console.log('原始返回：', data);
    }
  }
});
