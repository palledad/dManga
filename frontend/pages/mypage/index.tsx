import Segment from '../../components/Segment'
import Header from '../../components/Header'
import {
  Box,
  Grid,
  GridItem,
  Tabs,
  TabList,
  TabPanels,
  Tab,
  TabPanel,
  Button,
  Input,
} from '@chakra-ui/react'

const MyPage = () => {
  return (
    <>
      <Header></Header>
      <Box height={'20em'} background={'gray.100'}></Box>
      <Segment>
        <Box>ユーザー名</Box>
        <Box>wallet address</Box>
        <Box>自己紹介文</Box>
      </Segment>
      <Segment>
        <Tabs variant='enclosed'>
          <TabList>
            <Tab>購入した記事</Tab>
            <Tab>自分の記事</Tab>
            <Tab>持っている収益NFT</Tab>
          </TabList>

          <TabPanels>
            <TabPanel>
              <Input placeholder='検索' background={'gray.100'} />
              <Grid templateColumns='repeat(3, 1fr)'>
                <GridItem colSpan={1}>
                  <Segment>
                    記事サムネイル画像
                    <br />
                    タイトル
                    <br />
                    作者
                    <br />
                    作成日時
                    <br />
                    更新日時
                  </Segment>
                </GridItem>
              </Grid>
            </TabPanel>
            <TabPanel>
              <Input placeholder='検索' background={'gray.100'} />
              <Grid templateColumns='repeat(3, 1fr)'>
                <GridItem colSpan={1}>
                  <Segment>
                    タイトル
                    <br />
                    作成日時
                    <br />
                    更新日時
                    <br />
                    <Box display='flex' justifyContent='center' gap={2}>
                      <Button>編集</Button>
                      <Button>削除</Button>
                    </Box>
                  </Segment>
                </GridItem>
              </Grid>
            </TabPanel>
            <TabPanel>
              <Input placeholder='検索' background={'gray.100'} />
              <Grid templateColumns='repeat(3, 1fr)'>
                <GridItem colSpan={1}>
                  <Segment>
                    記事URL
                    <br />
                    記事タイトル
                    <br />
                    トークンID
                  </Segment>
                </GridItem>
              </Grid>
            </TabPanel>
          </TabPanels>
        </Tabs>
      </Segment>
    </>
  )
}

export default MyPage
