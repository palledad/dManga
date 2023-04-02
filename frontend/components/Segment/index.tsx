import { Box } from '@chakra-ui/react'

const Segment: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return (
    <Box borderRadius='sm' background='#FFFFFF' boxShadow='base' m={[2, 3]}>
      <Box p={5}>{children}</Box>
    </Box>
  )
}

export default Segment
