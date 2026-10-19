import axios, { AxiosError } from "axios";

const API_URL = "http://localhost:8080";

const api = axios.create({
  baseURL: API_URL,
  headers: {
    "Content-Type": "application/json",
  },
  timeout: 10000,
});


export type CreateProductRequest = {
  name: string;
  price: number;
  stock: number;
};

export type ProductResponse = {
  id: number;
  name: string;
  price: number;
  stock: number;
};

export type ApiError = {
  message: string;
  status?: number;
};

const handleApiError = (error: unknown): ApiError => {
  if (axios.isAxiosError(error)) {
    return {
      message:
        error.response?.data?.message ??
        error.message ??
        "Something went wrong",
      status: error.response?.status,
    };
  }

  return {
    message: error instanceof Error ? error.message : "Unknown error occurred",
  };
};

// ========================
// Product API
// ========================

class ProductClient {
  /**
   * Create a new product
   */
  async create(payload: CreateProductRequest): Promise<ProductResponse> {
    try {
      const { data } = await api.post<ProductResponse>("/products", payload);

      return data;
    } catch (error) {
      throw handleApiError(error);
    }
  }

  /**
   * Get product by ID
   */
  async getById(id: number): Promise<ProductResponse> {
    try {
      const { data } = await api.get<ProductResponse>(`/products/${id}`);

      return data;
    } catch (error) {
      throw handleApiError(error);
    }
  }

  /**
   * Get all products
   */
  async getAll(): Promise<ProductResponse[]> {
    try {
      const { data } = await api.get<ProductResponse[]>("/products");

      return data;
    } catch (error) {
      throw handleApiError(error);
    }
  }

  /**
   * Update an existing product
   */
  async update(
    id: number,
    payload: CreateProductRequest,
  ): Promise<ProductResponse> {
    try {
      const { data } = await api.put<ProductResponse>(
        `/products/${id}`,
        payload,
      );

      return data;
    } catch (error) {
      throw handleApiError(error);
    }
  }

  /**
   * Delete a product
   */
  async delete(id: number): Promise<ProductResponse> {
    try {
      const { data } = await api.delete<ProductResponse>(`/products/${id}`);

      return data;
    } catch (error) {
      throw handleApiError(error);
    }
  }
};


export const apiProductClient  = new  ProductClient()

