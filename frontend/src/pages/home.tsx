import { Tabs, TabsContent, TabsList, TabsTrigger } from "@cn/components/ui/tabs"
import { Card, CardContent, CardHeader, CardTitle} from "@cn/components/ui/card"
import {
    Carousel,
    CarouselContent,
    CarouselItem,
    CarouselNext,
    CarouselPrevious,
} from "@cn/components/ui/carousel"
import { Button } from "@cn/components/ui/button"
// import {
//     Select,
//     SelectContent,
//     SelectItem,
//     SelectTrigger,
//     SelectValue,
//   } from "@cn/components/ui/select"
  

const Home = () => {
    return (
        <div>
            <div className="flex flex-row-reverse">
                <Button variant="link">Membership</Button>
                <Button variant="link">Login</Button>
            </div>
            <Tabs defaultValue="Home" className="">
                <Card>
                    <TabsList>
                        <TabsTrigger value="Home">Home</TabsTrigger>
                        <TabsTrigger value="Bookticket">Book a bus ticket</TabsTrigger>
                        <TabsTrigger value="Contact">Contact</TabsTrigger>
                    </TabsList>
                </Card>
                <div className="flex justify-center">
                    <TabsContent value="Home">
                        <Carousel className="w-full max-w-m mt-24">
                            <CarouselContent>
                                {Array.from({ length: 5 }).map((_, index) => (
                                    <CarouselItem key={index}>
                                        <div className="p-1">
                                            <Card>
                                                <CardContent className="flex aspect-square items-center justify-center p-6">
                                                    <span className="text-4xl font-semibold">{index + 1}</span>
                                                </CardContent>
                                            </Card>
                                        </div>
                                    </CarouselItem>
                                ))}
                            </CarouselContent>
                            <CarouselPrevious />
                            <CarouselNext />
                        </Carousel>
                    </TabsContent>
                </div>
                <div>
                    <TabsContent value="Bookticket">
                        <div>
                            <Card className="m-10 p-4">
                                <CardHeader><CardTitle>Departure</CardTitle></CardHeader>
                                <CardContent>
                                    123456
                                </CardContent>
                            </Card>
                        </div>
                        <div>
                            <Card className="m-10 p-4">
                                <CardHeader><CardTitle>Return Departure</CardTitle></CardHeader>
                                <CardContent>
                                    123456
                                </CardContent>
                            </Card>
                        </div>
                        <Button variant="outline" className="ml-10">Search</Button>
                    </TabsContent>
                </div>
                <TabsContent value="Contact">
                    <div className="m-5">
                        <span className="text-xl font-semibold m-5">ติดต่อเรา</span>
                        <span></span>
                    </div>
                </TabsContent>
            </Tabs>
        </div>
    );
};
export default Home;