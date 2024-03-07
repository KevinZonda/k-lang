import {Button, Center, Container, Flex, Heading, Image, Link, Text, VStack, Wrap, WrapItem} from "@chakra-ui/react";
import {ArrowForwardIcon, DownloadIcon} from "@chakra-ui/icons";
import {useState} from "react";
import {useNavigate} from "react-router-dom";

export const highlightColor = '#3182ce';

export function NavBar() {
    return <Flex padding={'50px'} flexDirection={'row-reverse'} paddingBottom={0}>
        <Link color='black' href='https://docs.k-lang.org' _hover={{color: highlightColor, textDecoration: 'underline'}}>Docs</Link>
    </Flex>
}

export function Header() {
  const [count, setCount] = useState(0)
  const nav = useNavigate()
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
          <Text fontSize='xl' textAlign={'center'} color={'rgba(0, 0, 0, 0.65)'}>
            K Language is designed for beginner, aims to remove the barrier of learning programming.
          </Text>
        </Center>
      </Container>
      <Container>
        <Wrap spacing='30px' justify='center'>
          <WrapItem>
            <Button border={0} size='lg' colorScheme='blue' mt='24px' rightIcon={<ArrowForwardIcon/>}  onClick={() => window.open('https://playground.k-lang.org', "_blank")}>
              Try it now
            </Button>
          </WrapItem>
          <WrapItem>
            <Button border={0} size='lg' colorScheme='gray' mt='24px' onClick={() => nav('/download')} rightIcon={<DownloadIcon/>}>
              Download
            </Button>
          </WrapItem>

        </Wrap>

      </Container>
    </VStack>
  )
}