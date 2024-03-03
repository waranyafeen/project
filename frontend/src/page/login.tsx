import { Card, CardTitle, CardHeader, CardContent } from "@/components/ui/card";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { useEffect } from "react";
import { z } from "zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { toast } from "@/components/ui/use-toast";
import NavBar from "@/layout/narBar/navbar";
import { useLocation, useNavigate } from "react-router-dom";
import { User,  } from "@/interfaces";
import { http } from "@/services/index";

export type Role = "user" | "employee";

type Login = {
  Email: string;
  Password: string;
};

const Login = ({ role }: { role: Role }) => {
  const { state } = useLocation();
  const navigate = useNavigate();

  const LoginSchema = z.object({
    email: z
      .string({ required_error: "Email is required" })
      .email("Please enter a valid email"),
    password: z
      .string()
      .min(2, "Password must be at least 8 characters")
      .max(20, "Password must be at most 20 characters"),
  })

  const form = useForm<z.infer<typeof LoginSchema>>({
    resolver: zodResolver(LoginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  })

  const handleLogin = async (role: Role, data?: Login) => {
    const res = await http.Post<User>("/login/" + role, data || {});
    if (res.ok) {
      notifySuccess("Login Successful");
      navigate(state?.from || "/", { replace: true });
    } else {
      notifyError(res.error);
    }
  };

  function notifyError(msg: string) {
    if (msg === "Unauthorized") {
      return;
    }
    toast({
      title: msg === "record not found" ? "Account not found" : msg,
      duration: 3000,
      variant: "destructive",
    });
  }

  function notifySuccess(msg: string) {
    toast({
      title: msg,
      duration: 3000,
      variant: "success",
    });
  }

  async function onLogin() {
    const formData = form.getValues(); // รับค่าข้อมูลจากฟอร์ม
    const loginData: Login = {
      Email: formData.email,
      Password: formData.password
    };
    await handleLogin(role, loginData); // ส่งข้อมูลไปยัง handleLogin
  }
  
  useEffect(() => {
    onLogin();
  }, []);

  return (
    <div>
      <NavBar />
      <div className="w-full h-screen bg-secondary flex flex-col justify-center items-center gap-2 relative">
        <Card className="flex flex-col p-1 w-1/4 bg-violet-100">
          <CardHeader>
            <CardTitle className="text-lg font-bold text-primary text-center ">
              <span className="text-xl font-black text-primary mb-4">LOG IN</span>
            </CardTitle>
          </CardHeader>
          <CardContent>
            <Form {...form}>
              <form
                onSubmit={onLogin}
                className="flex flex-col gap-4"
              >
                <FormField
                  control={form.control}
                  name="email"
                  render={({ field }) => (
                    <FormItem>
                      <FormControl>
                        <Input placeholder="Username" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="password"
                  render={({ field }) => (
                    <FormItem>
                      <FormControl>
                        <Input placeholder="Password" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <Button type="submit" variant="success" className="flex flex-col items-center">Login</Button>
                <Button variant="link">Forget Password?</Button>
              </form>
            </Form>
          </CardContent>
        </Card>
      </div>
    </div>
  );
};
export default Login;
