import { useEffect, useState } from "react";
import { Gender } from "../../interfaces";
import { http } from "../../services/http";
import Form from "@shadcn/simplify/form";
import { useToast } from "@shadcn/ui/use-toast";
import { ToItemList } from "@src/utils";
import { z } from "zod";
import { Separator } from "@cn/components/ui/separator"

const UserCreate = () => {
  const { toast } = useToast();
  const [genders, setGender] = useState<Gender[]>([]);

  const userFormSchema = z.object({
    FirstName: z.string().min(2, "FirstName must be at least 2 characters"),
    LastName: z.string().min(2, "LastName must be at least 2 characters"),    Email: z.string().email({ message: "Invalid email address" }),
    Password: z.string().min(8, "Password must be at least 8 characters"),
    Phone: z.string().length(10, "Phone number must be 10 characters"),
    GenderID: z.number({ required_error: "Please select gender" }),
  });
  
  async function fetchGender() {
    const res = await http.Get<Gender[]>("/users/genders");
    if (res.ok) {
      setGender(res.data);
    }
  }
  useEffect(() => {
    fetchGender();
  }, []);

  async function onValid(formData: z.infer<typeof userFormSchema>) {
    const res = await http.Post<string>("/users", formData);
    if (res.ok) {
      toast({
        title: res.data,
        duration: 1500,
        variant: "success",
      });
    } else {
      toast({
        title: res.error,
        duration: 1500,
        variant: "destructive",
      });
    }
  }

  return (
    <Form
      className="w-full px-8 py-4 lg:px-16 h-with-nav flex flex-col gap-4"
      validator={userFormSchema}
      onValid={onValid}
      onInvalid={(errorFields) => console.log(errorFields)}
      fields={({ form, errors }) => (
        <>
          <h3 className="text-2xl font-medium lg:col-span-2 text-center">
            Personal Information
          </h3>
          <Separator className="lg:col-span-2 " />
          <div className="grid lg:grid-cols-2 p-4 gap-y-4">
            <div className="w-full h-full flex flex-col gap-3 max-w-lg justify-self-center">
              <Form.Label>First Name</Form.Label>
              <Form.Input useForm={form} name="FirstName" type="text" />
              <Form.Error field={errors.FirstName}></Form.Error>
              <Form.Label>Last Name</Form.Label>
              <Form.Input useForm={form} name="LastName" type="text" />
              <Form.Error field={errors.LastName}></Form.Error>
              <Form.Label>Gender</Form.Label>
              {genders.length > 0 ? (
                <Form.Select
                  valueAsNumber
                  useForm={form}
                  name="GenderID"
                  placeholder="Gender"
                  items={ToItemList(genders)}
                />
              ) : null}
              <Form.Error field={errors.GenderID}></Form.Error>
            </div>
            <div className="w-full h-full flex flex-col gap-3 max-w-lg justify-self-center">
              <Form.Label>Email</Form.Label>
              <Form.Input useForm={form} name="Email" type="text" />
              <Form.Error field={errors.Email}></Form.Error>
              <Form.Label>Password</Form.Label>
              <Form.Input
                useForm={form}
                name="Password"
                type="password"
                placeholder="********"
              />
              <Form.Error field={errors.Password}></Form.Error>
              <Form.Label>Phone Number</Form.Label>
              <Form.Input
                useForm={form}
                name="Phone"
                type="text"
                placeholder="0XXXXXXXXX"
              />
              <Form.Error field={errors.Phone}></Form.Error>
            </div>
            <Form.SubmitButton
              useForm={form}
              className="justify-self-center w-full max-w-lg lg:col-start-2"
            >
              Create
            </Form.SubmitButton>
          </div>
        </>
      )}
    />
  );
};
export default UserCreate;
