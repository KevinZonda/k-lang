import {Button, Center, Container, Heading, Image, Text, VStack, Wrap, WrapItem} from "@chakra-ui/react";
import {ArrowForwardIcon, DownloadIcon} from "@chakra-ui/icons";
import {useState} from "react";

export const highlightColor = '#3182ce';

export function Header() {
  const [count, setCount] = useState(0)
  return (
    <VStack paddingTop={'100px'} paddingBottom={'100px'}>
      <Center>
        {/*<Image w={32} h={32} src={"icon_512x512.png"}></Image>*/}
        <Image onClick={() => setCount(count + 1)} maxH={32}
               src={count % 2 === 0 ? "icon_512x512.png" : "K_Go.png"}/>
      </Center>
      <Container maxW='container.md'>
        <Center>
          <Heading mb={4} textAlign={'center'}>K Language: A
            <span style={{color: highlightColor}}> Modern Progressive </span>
            Teaching Language</Heading>
        </Center>
      </Container>
      <Container maxW='container.lg'>
        <Center>
          <Text fontSize='xl' textAlign={'center'}>
            K Language is designed for beginner, aims to remove the barrier of learning programming.
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