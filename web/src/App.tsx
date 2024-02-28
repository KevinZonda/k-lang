import {
  Text,
  VStack,
} from "@chakra-ui/react";
import {Header} from "./Head.tsx";
import {NumberSection} from "./numberSection.tsx";
import {FeatureSection} from "./Feature.tsx";


function App() {
  return <div style={{width: '100vw', minHeight: '100vh'}}>
    <VStack
      align='stretch'
      spacing={0}
    >
      <Header/>

      <FeatureSection/>

      <NumberSection/>

      <Footer/>

    </VStack>
  </div>
}

function Footer() {
  return <div style={{backgroundColor: '#EDF2F7', padding: '8px'}}>
    <Text textAlign={'center'}>
      K Language, build with Go, Antlr 4, powered by WebAssembly
    </Text>
    <Text textAlign={'center'}>
      Copyright &copy; 2024 Xiang Shi (xxs166@student.bham.ac.uk)
    </Text>
  </div>
}


export default App
