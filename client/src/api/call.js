import api from "./axios";
export const MakeApiCall = async (url, method, data) => {
  try {
    const response = await api.request({ url, method, data });
    return handleAxiosResponse(response);
  } catch (error) {
    return {
      success: false,
      error: error.response?.data?.error || "Something went wrong",
    };
  }
};

const handleAxiosResponse = (response) => {
  if (response?.data?.success) {
    return {
      success: true,
      data: response.data?.data,
      message: response.data?.message,
    };
  }
  return {
    success: false,
    error: response.data?.error || "Something went wrong",
  };
};
