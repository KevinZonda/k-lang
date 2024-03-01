import {highlightColor} from "./Head.tsx";
import {
  Badge,
  Box,
  Container,
  Heading,
  Link,
  ListItem,
  SimpleGrid,
  Text,
  UnorderedList
} from "@chakra-ui/react";

interface DownloadProp {
  url: string
  platform: string
  badge: string
  arch?: string
  desc?: string
}

const featuredDownload: DownloadProp[] = [
  {
    url: '',
    platform: 'Microsoft Windows',
    badge: 'Classic',
    arch: 'x64',
  }, {
    url: '',
    platform: 'Apple macOS',
    badge: 'Classic',
    arch: 'arm64',
  }, {
    url: '',
    platform: 'Apple macOS',
    badge: 'Classic',
    arch: 'x86_64',
  }, {
    url: '',
    platform: 'Linux',
    badge: 'Classic',
    arch: 'x86_64',
  }
]

function DownloadBox({url, platform, badge, desc, arch}: DownloadProp) {
  return <Box bg='rgb(237, 242, 247)' p={4} flexGrow={1} flexShrink={1} flexBasis={0}>
    <Text as={'b'}>{platform} {arch && " (" + arch + ")"} <Badge colorScheme='green'>{badge}</Badge></Text>
    <Text>{desc}</Text>
    <Link href={url}>Download</Link>
  </Box>
}

export function BarHeader({title}: { title?: string }) {
  return <div style={{width: '100vw'}}>
    <div style={{
      display: 'flex', alignItems: 'center',
      backgroundColor: highlightColor,
      paddingTop: '8px',
      paddingBottom: '8px'
    }}>
      <a href={'/'} style={{display: 'inline', fontSize: '32px', marginBlock: 0, color: 'white',
        paddingLeft: '16px', fontFamily: '-apple-system,BlinkMacSystemFont,"Segoe UI",Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Segoe UI Symbol"'}}>
        {title && title !== "" ? "K Language: " + title : "K Language"}
      </a>
    </div>
  </div>
}

export function DownloadPage() {
  return <div style={{width: '100vw'}}>
    <BarHeader/>

    <Container maxW='container.lg'>
      <div style={{paddingTop: '32px', paddingBottom: '32px'}}>
        <Heading as={'h2'}>Releases</Heading>
        <Text paddingTop={'16px'}>K Language provides several kinds of releases for different needs:</Text>
        <UnorderedList>
          <ListItem><Text as={'b'}>Classic:</Text> Classic version of K Language contains all functions and utilities
            except IDLE.</ListItem>
          <ListItem><Text as={'b'}>Minimal:</Text> Minimal version of K Language contains only interpreter and code
            runner.</ListItem>
          <ListItem><Text as={'b'}>IDLE:</Text> IDLE only.</ListItem>
        </UnorderedList>
        <SimpleGrid minChildWidth={'200px'} spacing={4} paddingTop={'16px'}>
          {featuredDownload.map((item) => <DownloadBox key={item.platform} {...item}/>)}
        </SimpleGrid>
      </div>
    </Container>
  </div>
}