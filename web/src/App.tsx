import {
    Box,
    Button,
    Center,
    Container,
    Heading,
    Image,
    StackDivider,
    Text,
    VStack,
    Wrap,
    WrapItem
} from "@chakra-ui/react";
import {ArrowForwardIcon, DownloadIcon} from "@chakra-ui/icons";

function App() {
    return <div style={{width: '100vw', minHeight: '100vh'}}>
        <VStack
            divider={<StackDivider borderColor='gray.200' width={'100%'} />}
            spacing={4}
            align='stretch'
        >
            <Header/>

            <Box/>

            <Footer/>

        </VStack>
    </div>
}

function Footer() {
    return <div>
        <Text textAlign={'center'}>
            K Language, build with Go, Antlr 4, powered by WebAssembly
        </Text>
        <Text textAlign={'center'}>
            Copyright &copy; 2024 Xiang Shi (xxs166@student.bham.ac.uk)
        </Text>
    </div>
}

function Header() {
  return (
        <VStack paddingTop={'100px'} paddingBottom={'100px'}>
            <Center>
                <Image w={32} h={32} src={"icon_512x512.png"}></Image>
            </Center>
            <Container maxW='container.md'>
                <Center>
                    <Heading mb={4}  textAlign={'center'}>K Language: A
                        <span style={{color: '#3182ce'}}> Modern Progressive </span>
                        Teaching Language</Heading>
                </Center>
            </Container>
            <Container maxW='container.lg'>
                <Center>
                    <Text fontSize='xl' textAlign={'center'}>
                        K Language is designed for beginner.
                    </Text>
                </Center>
            </Container>
            <Container>
                <Wrap spacing='30px' justify='center'>
                    <WrapItem>
                        <Button size='lg' colorScheme='blue' mt='24px' rightIcon={<ArrowForwardIcon/>}>
                            Try it now
                        </Button>
                    </WrapItem>
                    <WrapItem>
                        <Button size='lg' colorScheme='gray' mt='24px' rightIcon={<DownloadIcon/>}>
                            Download
                        </Button>
                    </WrapItem>

                </Wrap>

            </Container>
        </VStack>
  )
}

export default App
