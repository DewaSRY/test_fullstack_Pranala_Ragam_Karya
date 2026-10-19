import { useQuery } from "@tanstack/react-query";
import { apiProductClient, CreateProductRequest, ProductResponse } from "@/features/api";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export const productKeys = {
  all: ["products"] as const,
  lists: () => [...productKeys.all, "list"] as const,
  details: () => [...productKeys.all, "detail"] as const,
  detail: (id: number) => [...productKeys.details(), id] as const,
};

// Get all products
export const useProducts = () => {
  return useQuery({
    queryKey: productKeys.lists(),
    queryFn: apiProductClient.getAll,
  });
};

// Get product by ID
export const useProduct = (id: number) => {
  return useQuery({
    queryKey: productKeys.detail(id),
    queryFn: () => apiProductClient.getById(id),
    enabled: id > 0,
  });
};







// Create product
export const useCreateProduct = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateProductRequest) =>
      apiProductClient.create(payload),

    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: productKeys.all,
      });
    },
  });
};

// Update product
export const useUpdateProduct = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({
      id,
      payload,
    }: {
      id: number;
      payload: CreateProductRequest;
    }) => apiProductClient.update(id, payload),

    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({
        queryKey: productKeys.lists(),
      });

      queryClient.invalidateQueries({
        queryKey: productKeys.detail(variables.id),
      });
    },
  });
};

// Delete product
export const useDeleteProduct = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number) => apiProductClient.delete(id),

    onSuccess: (_, id) => {
      queryClient.invalidateQueries({
        queryKey: productKeys.all,
      });

      queryClient.removeQueries({
        queryKey: productKeys.detail(id),
      });
    },
  });
};
