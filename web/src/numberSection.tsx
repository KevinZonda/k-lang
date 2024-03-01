import {Box, Container, Heading, SimpleGrid, Text} from "@chakra-ui/react";
import {highlightColor} from "./Head.tsx";

export interface SimpleBlockProps {
  title: string
  content: string
}

export function NumberBlock({title, content}: SimpleBlockProps) {
  return <Box borderLeft={'2px solid white'} w={'100%'} minH={'50px'}>
    <Box style={{paddingLeft: '2em'}}>
      <Heading size={'3xl'} paddingBottom={'8px'}>{title}</Heading>
      <Text color={'rgba(255, 255, 255, 0.85)'}>{content}</Text>
    </Box>

  </Box>
}

export function NumberSection() {
  return <div style={{paddingTop: '50px', paddingBottom: '100px', backgroundColor: highlightColor, color: 'white'}}>
    <Heading textAlign={'center'} paddingBottom={'50px'}>Simple</Heading>
    <Container maxW='container.sm'>
      <SimpleGrid
        w='100%'
        templateColumns='repeat(auto-fill, minmax(200px, 1fr))'
        gap={8}
      >
        <NumberBlock title={'< 10 MiB'} content={'Core K Language Size'}/>
        <NumberBlock title={'22'} content={'K Language Keywords Count'}/>
        <NumberBlock title={'> 70%'} content={'Test Coverage'}/>
      </SimpleGrid>
    </Container>

  </div>
}