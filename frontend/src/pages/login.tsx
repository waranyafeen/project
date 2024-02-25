import { z } from "zod";
import { useAuth } from "@src/providers/auth";
import { useEffect } from "react";

const validLogin = z.object({
    Email: z
        .string({ required_error: "Email is required" })
        .email("Please enter a valid email"),
    Password: z
        .string()
        .min(2, "Password must be at least 8 characters long")
        .max(20, "Password must be at most 20 characters long"),
});

type TLogin = z.infer<typeof validLogin>;

const Login = () => {
    const { handleLogin } = useAuth();

    async function onLogin(data?: TLogin) {
        await handleLogin(role, data);
    }

    useEffect(() => {
        onLogin();
    }, []);
    
    return (
        <div>

        </div>
    );
};
export default Login;