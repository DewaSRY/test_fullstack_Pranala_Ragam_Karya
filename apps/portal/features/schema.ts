
import {z} from "zod"


export const CreateProductSchema = z.object({
    name: z.string(), 
    price: z.number().min(0), 
    stoc: z.number().min(0)
})

export type CreateProduc = z.infer<typeof CreateProductSchema>