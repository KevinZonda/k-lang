import {highlightColor} from "./Head.tsx";
import {
  Badge,
  Box,
  Container,
  Heading,
  Link,
  ListItem,
  SimpleGrid, Table, TableContainer, Tbody, Td,
  Text, Th, Thead, Tr,
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
    badge: 'Standard',
    arch: 'x64',
  }, {
    url: '',
    platform: 'Apple macOS',
    badge: 'Standard',
    arch: 'arm64',
  }, {
    url: '',
    platform: 'Apple macOS',
    badge: 'Standard',
    arch: 'x86_64',
  }, {
    url: '',
    platform: 'Linux',
    badge: 'Standard',
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
          <ListItem><Text as={'b'}>Standard:</Text> Classic version of K Language contains all functions and utilities
            except IDLE.</ListItem>
          <ListItem><Text as={'b'}>Minimal:</Text> Minimal version of K Language contains only interpreter and code
            runner.</ListItem>
          <ListItem><Text as={'b'}>IDLE:</Text> IDLE only.</ListItem>
        </UnorderedList>
        <SimpleGrid minChildWidth={'200px'} spacing={4} paddingTop={'16px'}>
          {featuredDownload.map((item) => <DownloadBox key={item.platform} {...item}/>)}
        </SimpleGrid>
      </div>

      <DownloadTable/>
    </Container>
  </div>
}

interface DlItem {
  Platform: string
  Arch: string
  Kind: string
  Hash: string
  Url: string
}

const dlT : DlItem[] = [
  {
    Platform: 'w',
    Arch: '64',
    Kind: 's',
    Hash: '',
    Url: ''
  }, {
  Platform: 'm',
    Arch: '64',
    Kind: 's',
    Hash: '',
    Url: ''
  }, {
    Platform: 'm',
    Arch: 'a64',
    Kind: 's',
    Hash: '',
    Url: ''
  }
]

function DownloadTable() {
  return <TableContainer w={'100%'}>
    <Table variant='simple'>
      <Thead>
        <Tr>
          <Th>Platform</Th>
          <Th>Architect</Th>
          <Th>Kind</Th>
          <Th>Hash</Th>
        </Tr>
      </Thead>
      <Tbody>
        {
          dlT.map((v) => {
            return <Tr>
              <Td>{v.Platform === 'w' ? 'Microsoft Windows' : v.Platform == 'l' ? 'Linux' : v.Platform == 'm' ? 'Apple macOS' : v.Platform }</Td>
              <Td>{v.Arch === '64' ? 'x64' : v.Arch === 'a64' ? 'arm64' : v.Arch}</Td>
              <Td><Badge colorScheme='green'>{v.Kind === 's' ? 'standard' : v.Kind === 'g' ? 'gui' : v.Kind === 'm' ? 'minimal' : v.Kind}</Badge></Td>
              <Td>{v.Hash === '' ? 'N/A' : v.Hash}</Td>

            </Tr>
          })
        }
      </Tbody>
    </Table>
  </TableContainer>
}