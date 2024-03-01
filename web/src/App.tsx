import {
  Box, Center,
  Text,
  VStack,
} from "@chakra-ui/react";
import {Header} from "./Head.tsx";
import {NumberSection} from "./numberSection.tsx";
import {FeatureSection} from "./Feature.tsx";
import SyntaxHighlighter from 'react-syntax-highlighter';
import {dracula} from 'react-syntax-highlighter/dist/esm/styles/hljs';

import klang from "./highlighter/klang.ts";
import {Light as SyntaxHighlighterL} from 'react-syntax-highlighter';

SyntaxHighlighterL.registerLanguage('klang', klang);


function App() {

  return <div style={{width: '100vw', minHeight: '100vh'}}>
    <VStack
      align='stretch'
      spacing={0}
    >
      <Header/>

      <SampleCode/>

      <FeatureSection/>

      <NumberSection/>

      <Footer/>

    </VStack>
  </div>
}

function Footer() {
  return <div style={{backgroundColor: 'rgb(237, 242, 247)', padding: '8px'}}>
    <Text textAlign={'center'}>
      K Language, build with Go, Antlr 4, powered by WebAssembly
    </Text>
    <Text textAlign={'center'}>
      Copyright &copy; 2024 Xiang Shi (xxs166@student.bham.ac.uk)
    </Text>
  </div>
}

function SampleCode() {
  return <Box backgroundColor={'rgb(40, 42, 54)'} paddingTop={'50px'} paddingBottom={'50px'}>
    <Center maxW='container.md'>
      <SyntaxHighlighter language="klang" style={dracula}>
        {'# Simple function\n' +
          'fn helloWorld(){\n' +
          '    return "Hello World from K Language!"\n' +
          '}\n' +
          'println(helloWorld())'}
      </SyntaxHighlighter>
    </Center>
  </Box>
}


export default App
