import {Button, Center, Container, Heading, Image, Text, VStack, Wrap, WrapItem} from "@chakra-ui/react";
import {ArrowForwardIcon, DownloadIcon} from "@chakra-ui/icons";
export const highlightColor = '#3182ce';
export function Header() {
  return (
    <VStack paddingTop={'100px'} paddingBottom={'100px'}>
      <Center>
        <Image w={32} h={32} src={"icon_512x512.png"}></Image>
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