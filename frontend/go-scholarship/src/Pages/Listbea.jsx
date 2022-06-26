import React from "react";
import {
  Box,
  Heading,
  Link,
  Image,
  Text,
  HStack,
  Tag,
  Container,
} from "@chakra-ui/react";
import { useScholarship } from "../Database/useScholarship";
import { useEffect } from "react";

const BlogTags = (props) => {
  return (
    <HStack spacing={2} marginTop={props.marginTop}>
      {props.tags.map((tag) => {
        return (
          <Tag size={"md"} variant="solid" colorScheme="orange" key={tag}>
            {tag}
          </Tag>
        );
      })}
    </HStack>
  );
};

export const BlogAuthor = (props) => {
  return (
    <HStack marginTop="2" spacing="2" display="flex" alignItems="center">
      <Image
        borderRadius="full"
        boxSize="40px"
        src={`${props.image}`}
        alt={`Avatar of ${props.name}`}
      />
      <Text fontWeight="medium">{props.name}</Text>
      <Text>—</Text>
      <Text>{props.date.toLocaleDateString()}</Text>
    </HStack>
  );
};

const ListBea = () => {
  const { scolarship, fetch } = useScholarship((state) => state);
  useEffect(() => {
    fetch();
  }, []);

  return (
    <Container maxW={"7xl"} p="12">
      <Heading as="h1"></Heading>
      {scolarship?.length > 0 ? (
        <>
          {scolarship?.map((item, index) => (
            <Box
              marginTop={{ base: "1", sm: "5" }}
              display="flex"
              flexDirection={{ base: "column", sm: "row" }}
              justifyContent="space-between"
              key={index}
            >
              <Box
                display="flex"
                flex="1"
                marginRight="3"
                position="relative"
                alignItems="center"
              >
                <Box
                  width={{ base: "100%", sm: "85%" }}
                  zIndex="2"
                  marginLeft={{ base: "0", sm: "5%" }}
                  marginTop="5%"
                >
                  <Link
                    textDecoration="none"
                    _hover={{ textDecoration: "none" }}
                  >
                    <Image
                      borderRadius="lg"
                      src={
                        item?.image.includes("http") ? item?.image : "https://images.unsplash.com/photo-1499951360447-b19be8fe80f5?ixlib=rb-1.2.1&ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&auto=format&fit=crop&w=800&q=80"
                      }
                      alt="some good alt text"
                      objectFit="contain"
                    />
                  </Link>
                </Box>
                <Box zIndex="1" width="100%" position="absolute" height="100%">
                  <Box
                    //   bgGradient={useColorModeValue(
                    //     'radial(orange.600 1px, transparent 1px)',
                    //     'radial(orange.300 1px, transparent 1px)'
                    //   )}
                    backgroundSize="20px 20px"
                    opacity="0.4"
                    height="100%"
                  />
                </Box>
              </Box>
              <Box
                display="flex"
                flex="1"
                flexDirection="column"
                justifyContent="center"
                marginTop={{ base: "3", sm: "0" }}
              >
                {
                  item?.category?.category_name && (<BlogTags tags={[item?.category?.category_name]} />)
                }
                
                <Heading marginTop="1" textAlign={"left"}>
                  <Link
                    textDecoration="none"
                    _hover={{ textDecoration: "none" }}
                  >
                    {item?.name}
                  </Link>
                </Heading>
                <Text
                  textAlign={"left"}
                  as="p"
                  marginTop="2"
                  color={"black.700"}
                  fontSize="lg"
                >
                  {
                    item?.description
                  }
                </Text>
                {
                  item?.user?.name && <BlogAuthor image={item?.image.includes("http") ? item?.image : "https://100k-faces.glitch.me/random-image"} name={item?.user?.name} date={new Date(item?.created_at)} />
                }
                
              </Box>
            </Box>
          ))}
        </>
      ) : (
        <Text>Loading...</Text>
      )}
    </Container>
  );
};

export default ListBea;
