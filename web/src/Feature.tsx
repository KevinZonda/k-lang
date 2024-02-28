import {
  Box,
  Card,
  CardBody,
  CardHeader,
  Container,
  Heading,
  SimpleGrid,
  Text
} from "@chakra-ui/react";
import {Icon} from "@chakra-ui/icons";
import {ReactElement} from "react";
import {highlightColor} from "./Head.tsx";
import {SiJupyter, SiGtk, SiVisualstudiocode} from "react-icons/si";
import { GiPlatform} from "react-icons/gi";
import {HiMiniCpuChip} from "react-icons/hi2";
import {PiCompassBold} from "react-icons/pi";

interface FeatureProp {
  title: string
  content: string
  icon: ReactElement
  iconColor: string
}

function FeatureCard({title, content, icon, iconColor}: FeatureProp) {
  return <Card>
    <div style={{padding: '1.25rem', paddingBottom: '0'}}>
      <Box display={'flex'} alignItems={'center'} textAlign={'center'} verticalAlign={'center'}
           justifyContent={'center'} backgroundColor={iconColor} borderRadius={'999'} color={'white'} fontSize={20}
           h={'2.5rem'} w={'2.5rem'}>
        {icon}
      </Box>
    </div>

    <CardHeader>
      <Heading size='md'>{title}</Heading>
    </CardHeader>
    <CardBody paddingTop={0}>
      <Text>{content}</Text>
    </CardBody>
  </Card>
}

export function FeatureSection() {
  return <div style={{paddingTop: '100px', paddingBottom: '100px', backgroundColor: '#EDF2F7'}}>
    <Heading textAlign={'center'} paddingBottom={'50px'}>Features</Heading>
    <Container maxW='container.md'>
      <SimpleGrid spacing={4} templateColumns='repeat(auto-fill, minmax(200px, 1fr))'>
        <FeatureCard
          iconColor={highlightColor}
          title={'Powerful'}
          content={'K Language' +
            ' is Turing Complete, means it\'s as powerful as other languages, i.e., Java and Python'}
          icon={<Icon as={HiMiniCpuChip}/>}/>
        <FeatureCard
          iconColor={highlightColor}
          title={'Legacy & Modern Patterns'}
          content={'K Languages supports C-like left type pattern but also modern right type pattern'}
          icon={<Icon as={PiCompassBold}/>}/>
        <FeatureCard
          iconColor={highlightColor}
          title={'Cross Platform'}
          content={'K Languages supports Windows, macOS, Linux and more. You even can run K Language on web.'}
          icon={<Icon as={GiPlatform}/>}/>
        <FeatureCard
          iconColor={highlightColor}
          title={'Editor Support'}
          content={'Code with your favorite text editor. K Language provides standard LSP & VSCode plugin.'}
          icon={<Icon as={SiVisualstudiocode}/>}/>
        <FeatureCard
          iconColor={highlightColor}
          title={'Built-in IDLE'}
          content={'Don\'t fancy a programming IDE? Use K Language as a built-in IDLE.'}
          icon={<Icon as={SiGtk}/>}/>
        <FeatureCard
          iconColor={highlightColor}
          title={'Jupyter Notebook'}
          content={'K Language as a Jupyter Kernel. Lecturing with interaction.'}
          icon={<Icon as={SiJupyter} />}/>
      </SimpleGrid>
    </Container>
  </div>
}