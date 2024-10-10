import axios from "axios";

const axiosInstance = axios.create({
  baseURL: "http://backend:3000",
});

export default axiosInstance;
