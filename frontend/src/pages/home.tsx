import { Card, CardContent } from "@shadcn/ui/card";
import NavBar from "@src/components/navbar/navBar";
import { Phone, Mail } from 'lucide-react';
import {
    Carousel,
    CarouselContent,
    CarouselItem,
    CarouselNext,
    CarouselPrevious,
} from "@cn/components/ui/carousel";
import Tour9 from "@src/assets/tour9.jpg"
import Tour8 from "@src/assets/tour8.jpg"
import Tour5 from "@src/assets/tour5.jpg"

const Home = () => {
    return (
        <main className="w-full h-screen flex flex-col items-center ">
            <NavBar />
            <Carousel className="w-screen max-w-lg mt-2">
                <CarouselContent>
                    <CarouselItem>
                        <div>
                            <Card className="w-full h-full">
                                <CardContent className="flex aspect-square items-center justify-center p-6">
                                    <img src={Tour5} alt="1" className="w-full h-full object-cover" />
                                </CardContent>
                            </Card>
                        </div>
                    </CarouselItem>
                    <CarouselItem>
                        <div>
                            <Card className="w-full h-full">
                                <CardContent className="flex aspect-square items-center justify-center p-6">
                                    <img src={Tour8} alt="Image 2" className="w-full h-full object-cover" />
                                </CardContent>
                            </Card>
                        </div>
                    </CarouselItem>
                    <CarouselItem>
                        <div>
                            <Card className="w-full h-full">
                                <CardContent className="flex aspect-square items-center justify-center p-6">
                                    <img src={Tour9} alt="Image 3" className="w-full h-full object-cover" />
                                </CardContent>
                            </Card>
                        </div>
                    </CarouselItem>
                </CarouselContent>
                <CarouselPrevious />
                <CarouselNext />
            </Carousel>

            <div className="mt-auto">
                <Card className="w-screen bg-gray-200">
                    <div className="text-xl mb-2 m-5">ติดต่อเรา</div>
                    <div className="text-m flex items-center m-5">
                        <Phone size={24} className="mr-2" />
                        <div>0971037598</div>
                    </div>
                    <div className="text-m flex items-center m-5">
                        <Mail size={24} className="mr-2" />
                        <div>B6400415@g.sut.ac.th</div>
                    </div>
                </Card>
            </div>
        </main>
    );
};

export default Home;
