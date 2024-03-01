import { z } from "zod";
import { Role, useAuth } from "@src/providers/auth";
import { useEffect } from "react";
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle
} from "@shadcn/ui/card";
import { Label } from "@shadcn/ui";
import { cn } from "@cn/utils";
import Form from "@cn/components/simplify/form";
import { Badge } from "@shadcn/ui/badge";

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

const Login = ({ role }: { role: Role }) => {
  const { handleLogin } = useAuth();

  async function onLogin(data?: TLogin) {
    await handleLogin(role, data);
  }

  useEffect(() => {
    onLogin();
  }, []);

  return (
    <div className="w-full h-screen bg-secondary flex flex-col justify-center items-center gap-2 relative">
      <p className="text-4xl font-black text-primary mb-4 animate-tracking-in-expand z-10 -mt-12">
        WARANRAT TOUR
      </p>
      <Card className="flex flex-col w-3/4 max-w-sm relative ">
        <CardHeader>
          <CardTitle className="text-lg font-bold text-primary text-center ">
            Log in
          </CardTitle>
        </CardHeader>
        <CardContent>
          {role !== "user" && (
            <Badge
              className={cn(
                "absolute top-2 right-2 rounded-full",
                role === "admin"
                  ? "bg-sky-500 hover:bg-sky-500/80"
                  : "bg-amber-500 hover:bg-amber-500/80"
              )}
            >
              {role}
            </Badge>
          )}
          <Form
            className="flex flex-col gap-4"
            validator={validLogin}
            onValid={onLogin}
            onInvalid={(errorFields) => console.log(errorFields)}
            fields={({ form, errors }) => (
              <>
                <Label>Email</Label>
                <Form.Input
                  useForm={form}
                  name="Email"
                  type="email"
                  placeholder="example@mail.com"
                />
                <Form.Error field={errors.Email} />
                <Label>Password</Label>
                <Form.Input
                  useForm={form}
                  name="Password"
                  type="password"
                  placeholder="********"
                />
                <Form.Error field={errors.Password} />
                <Form.SubmitButton useForm={form}>Log in</Form.SubmitButton>
              </>
            )}
          />
        </CardContent>
      </Card>
    </div>
  );
};
export default Login;