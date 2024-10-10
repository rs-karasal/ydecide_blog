import axios from "axios";

const axiosInstance = axios.create({
  baseURL: "http://194.32.141.224:3000",
});

export default axiosInstance;
