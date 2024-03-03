import { Card, CardContent } from "@/components/ui/card";
import NavBar from "@/layout/narBar/navbar";
import { Phone, Mail } from 'lucide-react';
import {
    Carousel,
    CarouselContent,
    CarouselItem,
    CarouselNext,
    CarouselPrevious,
} from "@/components/ui/carousel";
import Tour1 from "@/assets/tour1.jpg";
import Tour2 from "@/assets/tour2.jpg";
import Tour3 from "@/assets/tout3.jpg";

const Home = () => {
    return (
        <main className="flex flex-col h-screen justify-between">
            <NavBar />
            <div className="flex justify-center items-center flex-grow">
                <Carousel className="w-screen max-w-lg">
                    <CarouselContent>
                        <CarouselItem>
                            <div>
                                <Card className="w-full h-full">
                                    <CardContent className="flex aspect-square items-center justify-center p-6">
                                        <img src={Tour3} alt="1" className="w-full h-full object-cover" />
                                    </CardContent>
                                </Card>
                            </div>
                        </CarouselItem>
                        <CarouselItem>
                            <div>
                                <Card className="w-full h-full">
                                    <CardContent className="flex aspect-square items-center justify-center p-6">
                                        <img src={Tour2} alt="Image 2" className="w-full h-full object-cover" />
                                    </CardContent>
                                </Card>
                            </div>
                        </CarouselItem>
                        <CarouselItem>
                            <div>
                                <Card className="w-full h-full">
                                    <CardContent className="flex aspect-square items-center justify-center p-6">
                                        <img src={Tour1} alt="Image 3" className="w-full h-full object-cover" />
                                    </CardContent>
                                </Card>
                            </div>
                        </CarouselItem>
                    </CarouselContent>
                    <CarouselPrevious />
                    <CarouselNext />
                </Carousel>
            </div>
            <div className="w-screen">
                <div className="text-l mb-2 m-3">ติดต่อเรา</div>
                <div className="text-m flex items-center m-3">
                    <Phone size={18} className="mr-2" />
                    <div>0971037598</div>
                </div>
                <div className="text-m flex items-center m-3">
                    <Mail size={18} className="mr-2" />
                    <div>B6400415@g.sut.ac.th</div>
                </div>
            </div>
        </main>
    );
};

export default Home;